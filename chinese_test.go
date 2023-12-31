package chinese_s2t

import (
	"testing"
)

func TestT2s(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		{"2017年，A區15號# 人類終於回想起了曾經一度被他們所支配的恐怖，還有被囚禁於鳥籠中的那份屈辱", "2017年，A区15号# 人类终于回想起了曾经一度被他们所支配的恐怖，还有被囚禁于鸟笼中的那份屈辱"},
		{"我們被教導記住思想，而不是人，因為人可能失敗，他可能會被捕，他會被殺死，被遺忘，但400年後，思想仍可改變世界，我親眼目睹了，思想的威力，我見過人們以它為名殺戮，或是為了它獻出生命，但你不能親吻思想，也不能觸摸它，或擁抱它，思想不會流血，不會感到痛苦，它們沒有愛！", "我们被教导记住思想，而不是人，因为人可能失败，他可能会被捕，他会被杀死，被遗忘，但400年后，思想仍可改变世界，我亲眼目睹了，思想的威力，我见过人们以它为名杀戮，或是为了它献出生命，但你不能亲吻思想，也不能触摸它，或拥抱它，思想不会流血，不会感到痛苦，它们没有爱！"},
	}

	for _, test := range tests {
		output := T2s(test.input)
		if output != test.expected {
			t.Errorf("Test Failed: %v inputted, %v expected, received: %v", test.input, test.expected, output)
		}
	}
}

func TestS2t(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		{"2017年，A区15号# 人类终于回想起了曾经一度被他们所支配的恐怖，还有被囚禁于鸟笼中的那份屈辱", "2017年，A區15號# 人類終於回想起了曾經一度被他們所支配的恐怖，還有被囚禁於鳥籠中的那份屈辱"},
	}

	for _, test := range tests {
		output := S2t(test.input)
		if output != test.expected {
			t.Errorf("Test Failed: %v inputted, %v expected, received: %v", test.input, test.expected, output)
		}
	}
}
