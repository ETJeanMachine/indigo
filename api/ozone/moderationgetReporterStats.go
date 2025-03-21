// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package ozone

// schema: tools.ozone.moderation.getReporterStats

import (
	"context"

	"github.com/bluesky-social/indigo/xrpc"
)

// ModerationGetReporterStats_Output is the output of a tools.ozone.moderation.getReporterStats call.
type ModerationGetReporterStats_Output struct {
	Stats []*ModerationDefs_ReporterStats `json:"stats" cborgen:"stats"`
}

// ModerationGetReporterStats calls the XRPC method "tools.ozone.moderation.getReporterStats".
func ModerationGetReporterStats(ctx context.Context, c *xrpc.Client, dids []string) (*ModerationGetReporterStats_Output, error) {
	var out ModerationGetReporterStats_Output

	params := map[string]interface{}{
		"dids": dids,
	}
	if err := c.Do(ctx, xrpc.Query, "", "tools.ozone.moderation.getReporterStats", params, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
