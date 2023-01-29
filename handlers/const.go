package handlers

type PlagiarismResponse struct {
	Documents         []string    `json:"documents"`
	Sentences         []Sentence  `json:"sentences"`
	Paragraphs        []Paragraph `json:"paragraphs"`
	OverallSimilarity float32     `json:"overallsimilarityScore"`
}

type Sentence struct {
	SourceSentence   string  `json:"sourceSentence"`
	SourceID         int     `json:"sourceID"`
	ComparedSentence string  `json:"comparedSentence"`
	ComparedID       int     `json:"comparedID"`
	SimilarityScore  float64 `json:"similarityScore"`
}

type Paragraph struct {
	SourceParagraph   string  `json:"sourceParagraph"`
	SourceID          int     `json:"sourceID"`
	ComparedParagraph string  `json:"comparedParagraph"`
	ComparedID        int     `json:"comparedID"`
	SimilarityScore   float64 `json:"similarityScore"`
}

type PlagiarismRequest struct {
	Documents []string `json:"documents"`
}
