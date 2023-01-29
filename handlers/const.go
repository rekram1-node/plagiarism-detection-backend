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

var (
	mockResponse = []byte(`{"documents":["Normal science, the activity in which most scientists inevitably spend almost all their time, is predicated on the assumption that the scientific community knows what the world is like. Much of the success of the enterprise derives from the community's willingness to defend that assumption, if necessary at considerable cost. Normal science, for example, often suppresses fundamental novelties because they are necessarily subversive of its basic commitments (5).","Normal science, the activity in which most scientists inevitably spend almost all their time, is predicated on the assumption that the scientific community knows what the world is like. Some scientists say that the success of the enterprise comes from the community’s willingness to defend that assumption, if necessary at considerable cost. Normal science often suppresses fundamental novelties because they are necessarily subversive of its basic commitments."],"sentences":[{"sourceSentence":"Normal science, for example, often suppresses fundamental novelties because they are necessarily subversive of its basic commitments (5).","sourceID":0,"comparedSentence":"valid MLA","comparedID":0,"similarityScore":0},{"sourceSentence":"Normal science, the activity in which most scientists inevitably spend almost all their time, is predicated on the assumption that the scientific community knows what the world is like.","sourceID":0,"comparedSentence":"Normal science, the activity in which most scientists inevitably spend almost all their time, is predicated on the assumption that the scientific community knows what the world is like.","comparedID":1,"similarityScore":0.9999998807907104},{"sourceSentence":"Much of the success of the enterprise derives from the community's willingness to defend that assumption, if necessary at considerable cost.","sourceID":0,"comparedSentence":"Some scientists say that the success of the enterprise comes from the community’s willingness to defend that assumption, if necessary at considerable cost.","comparedID":1,"similarityScore":0.8998684883117676}],"paragraphs":[{"sourceParagraph":"Normal science, the activity in which most scientists inevitably spend almost all their time, is predicated on the assumption that the scientific community knows what the world is like. Much of the success of the enterprise derives from the community's willingness to defend that assumption, if necessary at considerable cost. ","sourceID":0,"comparedParagraph":"Normal science, the activity in which most scientists inevitably spend almost all their time, is predicated on the assumption that the scientific community knows what the world is like. Some scientists say that the success of the enterprise comes from the community’s willingness to defend that assumption, if necessary at considerable cost. Normal science often suppresses fundamental novelties because they are necessarily subversive of its basic commitments. ","comparedID":1,"similarityScore":0.9505704641342163}],"overallsimilarityScore":0.9505705}`)
)
