package openai

import (
	"encoding/json"
	"fmt"
	"strings"

	"forj-media-demo-backend/internal/models"
)

// simple helper to make a short summary out of raw notes
func summarize(raw string) string {
	raw = strings.TrimSpace(raw)
	if len(raw) == 0 {
		return "messy founder notes"
	}
	if len(raw) > 140 {
		return raw[:140] + "..."
	}
	return raw
}

func defaultIfEmpty(s, fallback string) string {
	if strings.TrimSpace(s) == "" {
		return fallback
	}
	return s
}

// Generate is now a LOCAL, rule-based generator.
// No external API, no keys, no billing.
func Generate(raw, product, audience, tone string) (string, error) {
	product = defaultIfEmpty(product, "B2B SaaS product")
	audience = defaultIfEmpty(audience, "B2B founders")
	tone = strings.ToLower(strings.TrimSpace(tone))
	if tone == "" {
		tone = "direct"
	}

	summary := summarize(raw)

	// crude tone variations
	var flavor string
	switch tone {
	case "story", "story-driven":
		flavor = "Tell a short story, keep it human and specific."
	case "contrarian":
		flavor = "Challenge a common belief and back it up with one clear example."
	case "mentor":
		flavor = "Speak calmly, like a trusted advisor, focusing on clarity instead of hype."
	default:
		flavor = "Be sharp and direct. Cut fluff, go straight to the point."
	}

	hooks := []string{
		fmt.Sprintf("The uncomfortable truth about how %s really buy from %s.", audience, product),
		fmt.Sprintf("Why your LinkedIn content isn’t converting %s (and what to fix this week).", audience),
		fmt.Sprintf("You don’t have a lead problem. You have a trust problem.", ),
		fmt.Sprintf("The moment I stopped ‘posting for the algorithm’ and started writing for actual %s.", audience),
		fmt.Sprintf("If your %s feed feels noisy, here’s how to become the signal.", audience),
	}

	postOutlines := []string{
		"Hook → short context from these notes → 1 painful problem → the shift in thinking → 3 bullet takeaways → soft CTA.",
		"Hook → describe what most teams are doing wrong → show 1 real example from your customer conversations → new approach → invite people to DM you.",
		"Hook → one-liner about trust vs volume → explain why more posting didn’t fix the problem → outline a simple weekly system → ask for feedback.",
	}

	fullPosts := []string{
		fmt.Sprintf(`Everyone says “post more” on LinkedIn.

But the founders I talk to aren’t short on content. They’re short on signal.

%s

That’s why I like to start from messy notes like these and ask:
- What pain keeps repeating?
- What decision is stuck?
- What’s the uncomfortable truth behind it?

From there, we turn chaos into a handful of sharp hooks and posts that actually move %s closer to a decision — not just more impressions.

You don’t need more noise. You need clearer stories about real problems.`, summary, audience),
		fmt.Sprintf(`You don’t have a lead problem. You have a trust problem.

If %s don’t feel like you understand their world, they’ll scroll past your content no matter how “optimized” it is.

Here’s a simple way to fix that:
1. Write down the last 3 real conversations you had with customers or prospects.
2. Pull out the exact phrases they used to describe the problem.
3. Turn each one into a post that:
   - starts with their words
   - shows what most people do wrong
   - and offers one small, concrete shift.

%s

That’s the core idea behind this intake engine: take real founder brain dumps and turn them into content that sounds like you — and lands with the right people.`, audience, flavor),
	}

	resp := models.GenerateResponse{
		Hooks:        hooks,
		PostOutlines: postOutlines,
		FullPosts:    fullPosts,
	}

	data, err := json.Marshal(resp)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
