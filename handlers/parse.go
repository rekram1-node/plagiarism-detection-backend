package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rekram1-node/text-processor/text"
)

func compare(docs []string, model *text.Word2Vec) (*PlagiarismResponse, error) {
	if len(docs) < 2 {
		return nil, fmt.Errorf("not enough documents")
	}

	document1, mlaCitations1 := citationCheck(docs[0])
	document2, _ := citationCheck(docs[1])

	plagiarismRes, err := constructResponse(document1, document2, mlaCitations1, model)
	if err != nil {
		return nil, err
	}

	plagiarismRes.Documents = docs

	return plagiarismRes, nil
}

func constructResponse(t1, t2 string, mlaCitations1 []string, model *text.Word2Vec) (*PlagiarismResponse, error) {
	t1Paragraphs, t1Sentences, err := text.ExtractAll(t1)
	if err != nil {
		return nil, err
	}

	t2Paragraphs, t2Sentences, err := text.ExtractAll(t2)
	if err != nil {
		return nil, err
	}

	sentenceMap, err := model.MostSimilarSentences(t1Sentences, t2Sentences)
	if err != nil {
		return nil, err
	}

	paragraphMap, err := model.MostSimilarParagraphs(t1Paragraphs, t2Paragraphs)
	if err != nil {
		return nil, err
	}

	overallSimiliarty, err := model.OverallSimilarity(t1, t2)
	if err != nil {
		return nil, err
	}

	response := &PlagiarismResponse{
		OverallSimilarity: overallSimiliarty,
	}

	response.AddMla(mlaCitations1)
	response.AddSentences(sentenceMap)
	response.AddParagraphs(paragraphMap)

	return response, nil
}

func (p *PlagiarismResponse) AddParagraphs(paragraphMap map[string]*text.ParaAndScore) {
	for key, element := range paragraphMap {
		p.Paragraphs = append(p.Paragraphs, Paragraph{
			SourceParagraph:   key,
			SourceID:          0,
			ComparedParagraph: element.Paragraph,
			ComparedID:        1,
			SimilarityScore:   float64(element.Score),
		})
	}
}

func (p *PlagiarismResponse) AddSentences(sentenceMap map[string]*text.SentAndScore) {
	for key, element := range sentenceMap {
		p.Sentences = append(p.Sentences, Sentence{
			SourceSentence:   key,
			SourceID:         0,
			ComparedSentence: element.Sentence,
			ComparedID:       1,
			SimilarityScore:  float64(element.Score),
		})
	}
}
func (p *PlagiarismResponse) AddMla(mlaCitations1 []string) {
	for _, mlacitation := range mlaCitations1 {
		p.Sentences = append(p.Sentences, Sentence{
			SourceSentence:   mlacitation,
			SourceID:         0,
			ComparedSentence: "valid MLA",
			ComparedID:       0,
			SimilarityScore:  float64(0.00),
		})
	}
}

func getDocumentsFromBody(w http.ResponseWriter, r *http.Request) ([]string, error) {
	var req PlagiarismRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, fmt.Errorf("failed to unmarshal body: %w", err)
	}

	return req.Documents, nil
}
