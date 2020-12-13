package twitter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChunkContent(t *testing.T) {
	twitClient := twitterClient{}
	testcases := map[string]struct {
		input  string
		output []string
	}{
		"Less than 280 chars": {
			input:  "One two three",
			output: []string{"One two three"},
		},
		"Two groups split on space": {
			input:  "One two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two three One two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two three",
			output: []string{"One two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two three One", "two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two three"},
		},
		"Three groups split on space": {
			input: "One two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two three",
			output: []string{
				"One two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne", "two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two", "threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two three"},
		},
		"Two groups split on new line": {
			input:  "One two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two three One\ntwo threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two three",
			output: []string{"One two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two three One", "two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two three"},
		},
		"Two groups split on tab": {
			input:  "One two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two three One\ttwo threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two three",
			output: []string{"One two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two three One", "two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two threeOne two three"},
		},
		"Multi byte runes": {
			input: `He aha koa. Hai te tokorima a Māui.

			It does not matter. I have the five of Māui.
			(If the host apologises for the lack of cutlery available, the guest replies that he has his fingers - the five of Māui.)`,
			output: []string{`He aha koa. Hai te tokorima a Māui.

			It does not matter. I have the five of Māui.
			(If the host apologises for the lack of cutlery available, the guest replies that he has his fingers - the five of Māui.)`},
		},
		"Tinirau": {
			input:  "Tutunui was a whale belonging to the chief Tinirau, who lent it to the tohunga Kae to convey him to his home. Kae killed the whale and his people cooked it, covering the flesh with leaves of the koromiko shrub. Tinirau looked in vain for the return of his whale, but when he smelled the roasting meat he knew what had happened and took revenge on Kae. When koromiko branches are thrown on the fire they give out a distinctive odour, which gave rise to this saying.",
			output: []string{"Tutunui was a whale belonging to the chief Tinirau, who lent it to the tohunga Kae to convey him to his home. Kae killed the whale and his people cooked it, covering the flesh with leaves of the koromiko shrub. Tinirau looked in vain for the return of his whale, but when he", "smelled the roasting meat he knew what had happened and took revenge on Kae. When koromiko branches are thrown on the fire they give out a distinctive odour, which gave rise to this saying."},
		},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			output := twitClient.chunkContent(tc.input)
			assert.Equal(t, len(tc.output), len(output), "Outputs have different lengths")
			for i := range tc.output {
				assert.Equalf(t, tc.output[i], output[i], "output strings do not match at position %d", i)
			}
		})
	}
}
