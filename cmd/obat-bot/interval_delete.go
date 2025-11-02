// OUR BRAINS ARE THINKING ‼️‼️
package main

import (
	"github.com/TBroz15/OUR-BRAINS-ARE-THINKING/internals/helpers"
	"github.com/TBroz15/OUR-BRAINS-ARE-THINKING/internals/helpers/yt"
	"github.com/charmbracelet/log"
)

/**
 * Use this binary to tries to delete new comments that
 * does not contain the words.
 *
 * Use this for 24/7 cleaning up
 */

func main() {
	ytService := yt.CreateYouTubeService()

	// Some reusable variables
	approvedIds := []string{}
	rejectedIds := []string{}
	bannedIds := []string{}
	violators := map[string]byte{}

	comments, err := ytService.
		CommentThreads.List([]string{"snippet"}).
		ModerationStatus("heldForReview").
		MaxResults(100).
		Order("time").
		// exclusive to my video only.
		// if you fork this repo and
		// you want to change it for your video, feel free!
		//
		// -tuxebro
		VideoId("Eov6cHQS6cM").
		Do()

	if err != nil {
		log.Fatal(err)
		return
	}

	for _, item := range comments.Items {
		comment := item.Snippet.TopLevelComment

		author := comment.Snippet.AuthorChannelId.Value
		text := comment.Snippet.TextDisplay

		isApproved := helpers.HasTheWords(text)

		if isApproved {
			approvedIds = append(approvedIds, comment.Id)
			continue
		}

		if _, ok := violators[author]; !ok {
			violators[author] = byte(0)
		}

		violators[author]++

		if violators[author] >= 3 {
			bannedIds = append(bannedIds, comment.Id)
			continue
		}

		rejectedIds = append(rejectedIds, comment.Id)
	}

	ytService.Comments.SetModerationStatus(approvedIds, "heldForReview").Do()
	ytService.Comments.SetModerationStatus(rejectedIds, "rejected").Do()
	ytService.Comments.SetModerationStatus(bannedIds, "rejected").BanAuthor(true).Do()

	approvedIds = helpers.ClearSlice(approvedIds)
	rejectedIds = helpers.ClearSlice(rejectedIds)
}
