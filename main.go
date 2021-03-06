// package whichlang is a suite of Machine Learning
// tools to classify programming languages.
package whichlang

import (
	"github.com/unixpickle/whichlang/gaussbayes"
	"github.com/unixpickle/whichlang/idtree"
	"github.com/unixpickle/whichlang/knn"
	"github.com/unixpickle/whichlang/neuralnet"
	"github.com/unixpickle/whichlang/svm"
	"github.com/unixpickle/whichlang/tokens"
)

type Classifier interface {
	// Classify classifies a tokenized source file.
	//
	// The returned string is the name of the
	// programming language in which the file is
	// most likely written in.
	Classify(tokens.Freqs) string

	// Languages returns all possible languages
	// that Classify() might return.
	// The result is not sorted, and its order
	// may change across calls.
	// Callers should not modify the returned slice.
	Languages() []string

	// Encode serializes this classifier as binary
	// data.
	Encode() []byte
}

// A Trainer generates a Classifier using
// a collection of tokenized sample files.
type Trainer func(map[string][]tokens.Freqs) Classifier

// A Decoder decodes a certain type of
// Classifier from binary data.
type Decoder func(d []byte) (Classifier, error)

// ClassifierNames is an array containing the
// names of every supported classifier.
var ClassifierNames = []string{"idtree", "neuralnet", "knn", "svm", "gaussbayes"}

// Trainers maps classifier names to their
// corresponding Trainers.
var Trainers = map[string]Trainer{
	"idtree": func(freqs map[string][]tokens.Freqs) Classifier {
		return idtree.Train(freqs)
	},
	"neuralnet": func(freqs map[string][]tokens.Freqs) Classifier {
		return neuralnet.Train(freqs)
	},
	"knn": func(freqs map[string][]tokens.Freqs) Classifier {
		return knn.Train(freqs)
	},
	"svm": func(freqs map[string][]tokens.Freqs) Classifier {
		return svm.Train(freqs)
	},
	"gaussbayes": func(freqs map[string][]tokens.Freqs) Classifier {
		return gaussbayes.Train(freqs)
	},
}

// Decoders maps classifier names to their
// corresponding Decoders.
var Decoders = map[string]Decoder{
	"idtree": func(d []byte) (Classifier, error) {
		return idtree.DecodeClassifier(d)
	},
	"neuralnet": func(d []byte) (Classifier, error) {
		return neuralnet.DecodeNetwork(d)
	},
	"knn": func(d []byte) (Classifier, error) {
		return knn.DecodeClassifier(d)
	},
	"svm": func(d []byte) (Classifier, error) {
		return svm.DecodeClassifier(d)
	},
	"gaussbayes": func(d []byte) (Classifier, error) {
		return gaussbayes.DecodeClassifier(d)
	},
}

// Descriptions maps classifier names to
// one-line descriptions of the classifier.
var Descriptions = map[string]string{
	"idtree":     "decision trees generated with ID3",
	"neuralnet":  "feedforward neural network",
	"knn":        "K-nearest neighbors",
	"svm":        "support vector machines",
	"gaussbayes": "naive Bayes with Gaussians",
}
