package rules

import (
	"fmt"

	"github.com/bluesky-social/indigo/automod"
	"github.com/bluesky-social/indigo/automod/countstore"
)

var interactionDailyThreshold = 800

var _ automod.RecordRuleFunc = InteractionChurnRule

// looks for accounts which do frequent interaction churn, such as follow-unfollow.
func InteractionChurnRule(c *automod.RecordContext) error {

	did := c.Account.Identity.DID.String()
	switch c.RecordOp.Collection {
	case "app.bsky.feed.like":
		c.Increment("like", did)
		created := c.GetCount("like", did, countstore.PeriodDay)
		deleted := c.GetCount("unlike", did, countstore.PeriodDay)
		ratio := float64(deleted) / float64(created)
		if created > interactionDailyThreshold && deleted > interactionDailyThreshold && ratio > 0.5 {
			c.Logger.Info("high-like-churn", "created-today", created, "deleted-today", deleted)
			c.AddAccountFlag("high-like-churn")
			c.ReportAccount(automod.ReportReasonSpam, fmt.Sprintf("interaction churn: %d likes, %d unlikes today (so far)", created, deleted))
			c.Notify("slack")
		}
	case "app.bsky.graph.follow":
		c.Increment("follow", did)
		created := c.GetCount("follow", did, countstore.PeriodDay)
		deleted := c.GetCount("unfollow", did, countstore.PeriodDay)
		ratio := float64(deleted) / float64(created)
		if created > interactionDailyThreshold && deleted > interactionDailyThreshold && ratio > 0.5 {
			c.Logger.Info("high-follow-churn", "created-today", created, "deleted-today", deleted)
			c.AddAccountFlag("high-follow-churn")
			c.ReportAccount(automod.ReportReasonSpam, fmt.Sprintf("interaction churn: %d follows, %d unfollows today (so far)", created, deleted))
			c.Notify("slack")
		}
		// just generic bulk following
		followRatio := float64(c.Account.FollowersCount) / float64(c.Account.FollowsCount)
		if created > interactionDailyThreshold && c.Account.FollowsCount > 2000 && followRatio < 0.2 {
			c.Logger.Info("bulk-follower", "created-today", created)
			c.AddAccountFlag("bulk-follower")
			c.ReportAccount(automod.ReportReasonSpam, fmt.Sprintf("bulk following: %d follows today (so far)", created))
			//c.Notify("slack")
		}
	}
	return nil
}

var _ automod.RecordRuleFunc = DeleteInteractionRule

func DeleteInteractionRule(c *automod.RecordContext) error {
	did := c.Account.Identity.DID.String()
	switch c.RecordOp.Collection {
	case "app.bsky.feed.like":
		c.Increment("unlike", did)
	case "app.bsky.graph.follow":
		c.Increment("unfollow", did)
	}
	return nil
}
