package handlers

type PlagiarismResponse struct {
	Documents []string `json:"documents"`
	Sentences []struct {
		SourceSentence   string  `json:"sourceSentence"`
		SourceID         int     `json:"sourceID"`
		ComparedSentence string  `json:"comparedSentence"`
		ComparedID       int     `json:"comparedID"`
		SimilarityScore  float64 `json:"similarityScore"`
	} `json:"sentences"`
	Paragraphs []struct {
		SourceParagraph   string  `json:"sourceParagraph"`
		SourceID          int     `json:"sourceID"`
		ComparedParagraph string  `json:"comparedParagraph"`
		ComparedID        int     `json:"comparedID"`
		SimilarityScore   float64 `json:"similarityScore"`
	} `json:"paragraphs"`
}
