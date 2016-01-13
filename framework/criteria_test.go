package framework

import (
	"os"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestCriteria(t *testing.T) {
	suite.Run(t, new(CriteriaSuite))
}

type CriteriaSuite struct {
	suite.Suite
	data []*ServiceInformation
}

func (suite *CriteriaSuite) SetupTest() {
	os.Clearenv()
	suite.data = []*ServiceInformation{}
	suite.data = append(suite.data,
		&ServiceInformation{
			ID:        "imagea-taga",
			ImageName: "imagea",
			ImageTag:  "taga",
			Instances: []*Instance{
				&Instance{
					ID:            "cida",
					Host:          "miniona",
					ContainerName: "ca",
					Status:        InstanceUp,
				},
				&Instance{
					ID:            "cidb",
					Host:          "miniona",
					ContainerName: "cb",
					Status:        InstanceUp,
				},
				&Instance{
					ID:            "cidc",
					Host:          "minionb",
					ContainerName: "ca",
					Status:        InstanceDown,
				},
			},
		})
	suite.data = append(suite.data,
		&ServiceInformation{
			ID:        "imagea-tagb",
			ImageName: "imagea",
			ImageTag:  "tagb",
			Instances: []*Instance{
				&Instance{
					ID:            "cide",
					Host:          "miniona",
					ContainerName: "cc",
					Status:        InstanceUp,
				},
				&Instance{
					ID:            "cidf",
					Host:          "miniona",
					ContainerName: "cd",
					Status:        InstanceDown,
				},
				&Instance{
					ID:            "cidg",
					Host:          "minionb",
					ContainerName: "cb",
					Status:        InstanceDown,
				},
			},
		})
	suite.data = append(suite.data,
		&ServiceInformation{
			ID:        "imageb-tagc",
			ImageName: "imageb",
			ImageTag:  "tagc",
			Instances: []*Instance{
				&Instance{
					ID:            "cidq",
					Host:          "minionc",
					ContainerName: "cq",
					Status:        InstanceUp,
				},
				&Instance{
					ID:            "cidw",
					Host:          "miniond",
					ContainerName: "cw",
					Status:        InstanceDown,
				},
				&Instance{
					ID:            "cide",
					Host:          "miniond",
					ContainerName: "ce",
					Status:        InstanceUp,
				},
			},
		})
}

func (suite *CriteriaSuite) assertLenCriteria(criteria ServiceInformationCriteria, length int) {
	result := criteria.MeetCriteria(suite.data)
	assert.Len(suite.T(), result, length)
}

func (suite *CriteriaSuite) assertImageNameAndImageTagRegexpCriteria(name string, length int) {
	suite.assertLenCriteria(&ImageNameAndImageTagRegexpCriteria{regexp.MustCompile(name)}, length)
}

func (suite *CriteriaSuite) TestImageNameAndImageTagRegexpCriteria() {
	suite.assertImageNameAndImageTagRegexpCriteria("imagea", 2)
}

func (suite *CriteriaSuite) TestAndCriteria() {
	suite.assertLenCriteria(&AndCriteria{
		&ImageNameAndImageTagRegexpCriteria{regexp.MustCompile("imagea")},
		&ImageNameAndImageTagRegexpCriteria{regexp.MustCompile("taga")},
	}, 1)
	suite.assertLenCriteria(&AndCriteria{
		&ImageNameAndImageTagRegexpCriteria{regexp.MustCompile("imagea")},
		&ImageNameAndImageTagRegexpCriteria{regexp.MustCompile("tagb")},
	}, 1)
	suite.assertLenCriteria(&AndCriteria{
		&ImageNameAndImageTagRegexpCriteria{regexp.MustCompile("imagea")},
		&ImageNameAndImageTagRegexpCriteria{regexp.MustCompile("imageb")},
	}, 0)
}

func (suite *CriteriaSuite) TestOrCriteria() {
	suite.assertLenCriteria(&OrCriteria{
		&ImageNameAndImageTagRegexpCriteria{regexp.MustCompile("imagea")},
		&ImageNameAndImageTagRegexpCriteria{regexp.MustCompile("imageb")},
	}, 3)
}
