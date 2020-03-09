package msg

import (
	"net/url"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Entity contains the type and the content of one entity.
type Entity struct {
	Type    string
	Content string
}

const (
	// EntityTypeURL is the identifier string for URL type of entity.
	EntityTypeURL = "url"
	// EntityTypeMention is the identifier string for mention type of entity.
	EntityTypeMention = "mention"
	// EntityTypeEmail is the identifier string for email type of entity.
	EntityTypeEmail = "email"
	// EntityTypeCode is the identifier string for code type of entity.
	EntityTypeCode = "code"
	// EntityTypePre is the identifier string for pre type of entity.
	EntityTypePre = "pre"
	// EntityTypeTextLink is the identifier string for text link type of entity.
	EntityTypeTextLink = "text_link"
	// EntityTypeTextMention is the identifier string for text mention type of entity.
	EntityTypeTextMention = "text_mention"
)

// GetEntities return all the entities in one message.
func GetEntities(m *tgbotapi.Message) []*Entity {
	entities := []*Entity{}
	for _, e := range *m.Entities {
		entityRaw := string([]rune(m.Text)[e.Offset : e.Offset+e.Length])
		switch e.Type {
		case EntityTypeURL:
			if _, err := url.ParseRequestURI(entityRaw); err != nil {
				entityRaw = "http://" + entityRaw
			}
			entities = append(entities, &Entity{
				Type:    EntityTypeURL,
				Content: entityRaw,
			})
		case EntityTypeMention:
			entities = append(entities, &Entity{
				Type:    EntityTypeMention,
				Content: "https://t.me/" + entityRaw[1:],
			})
		case EntityTypeEmail:
			entities = append(entities, &Entity{
				Type:    EntityTypeEmail,
				Content: entityRaw,
			})
		case EntityTypeCode:
			entities = append(entities, &Entity{
				Type:    EntityTypeCode,
				Content: entityRaw,
			})
		case EntityTypePre:
			entities = append(entities, &Entity{
				Type:    EntityTypePre,
				Content: entityRaw,
			})
		case EntityTypeTextLink:
			entities = append(entities, &Entity{
				Type:    EntityTypeTextLink,
				Content: e.URL,
			})
		case EntityTypeTextMention:
			entities = append(entities, &Entity{
				Type:    EntityTypeTextMention,
				Content: e.User.String(),
			})
		}
	}
	return entities
}
