package factories

import (
	"github.com/Gopherlinzy/gin-vue3-admin/app/models/topic"
	"github.com/bxcodec/faker/v4"
)

func MakeTopics(count int) []topic.Topic {

	var objs []topic.Topic

	// 设置唯一性，如 Topic 模型的某个字段需要唯一，即可取消注释
	// faker.SetGenerateUniqueValues(true)

	for i := 0; i < count; i++ {
		topicModel := topic.Topic{
			Title:      faker.Sentence(),
			Body:       faker.Paragraph(),
			UserID:     "1",
			CategoryID: "3",
		}
		objs = append(objs, topicModel)
	}

	return objs
}
