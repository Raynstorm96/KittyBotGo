package types

import (
	"strings"

	"github.com/DisgoOrg/disgo/core"
	"github.com/DisgoOrg/disgo/core/events"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func NewCommandMap(bot *Bot) *CommandMap {
	return &CommandMap{
		bot:      bot,
		commands: make(map[string]Command),
	}
}

type CommandMap struct {
	bot      *Bot
	commands map[string]Command
}

func (m *CommandMap) OnEvent(event core.Event) {
	if e, ok := event.(*events.ApplicationCommandInteractionEvent); ok {
		if cmd, ok := m.commands[e.Data.Name()]; ok {
			switch d := e.Data.(type) {
			case core.SlashCommandInteractionData:
				if cmd.CommandHandler != nil {
					if handler, ok := cmd.CommandHandler[buildCommandPath(d.SubCommandName, d.SubCommandGroupName)]; ok {
						if err := handler(m.bot, getMessagePrinter(e.BaseInteraction), e); err != nil {
							m.bot.Logger.Errorf("Failed to handle command \"%s\": %s", e.Data.Name(), err)
						}
						return
					}
				}
				m.bot.Logger.Errorf("No command handler for \"%s\"", e.Data.Name())
			}
		}
	} else if e, ok := event.(*events.AutocompleteInteractionEvent); ok {
		if cmd, ok := m.commands[e.Data.CommandName]; ok {
			if cmd.AutoCompleteHandler != nil {
				if handler, ok := cmd.AutoCompleteHandler[buildCommandPath(e.Data.SubCommandName, e.Data.SubCommandGroupName)]; ok {
					if err := handler(m.bot, getMessagePrinter(e.BaseInteraction), e); err != nil {
						m.bot.Logger.Errorf("Failed to handle autocomplete for \"%s\": %s", e.Data.CommandName, err)
					}
					return
				}
			}
			m.bot.Logger.Errorf("No autocomplete handler for command \"%s\"", e.Data.CommandName)
		}
	} else if e, ok := event.(*events.ComponentInteractionEvent); ok {
		customID := e.Data.ID().String()
		if !strings.HasPrefix(customID, "cmd:") {
			return
		}
		args := strings.Split(customID, ":")
		cmdHandler, action := args[1], args[2]
		cmdName := strings.Split(cmdHandler, "/")[0]
		if cmd, ok := m.commands[cmdName]; ok {
			if cmd.ComponentHandler != nil {
				if handler, ok := cmd.ComponentHandler[cmdHandler]; ok {
					if err := handler(m.bot, getMessagePrinter(e.BaseInteraction), e, action); err != nil {
						m.bot.Logger.Errorf("Failed to handle component interaction for \"%s\": %s", cmdName, err)
					}
					return
				}
			}
			m.bot.Logger.Errorf("No component handler for action \"%s\" on command \"%s\"", action, cmdName)
		}
	}
}

func getMessagePrinter(i *core.BaseInteraction) *message.Printer {
	lang, err := language.Parse(i.Locale.Code())
	if err != nil && i.GuildLocale != nil {
		i.Bot.Logger.Info("Failed to parse locale code, falling back to guild locale")
		lang, _ = language.Parse(i.GuildLocale.Code())
	}
	if lang == language.Und {
		i.Bot.Logger.Info("Failed to parse locale code, falling back to default locale")
		lang = language.English
	}
	return message.NewPrinter(lang)
}

func buildCommandPath(subcommand *string, subcommandGroup *string) string {
	var path string
	if subcommand != nil {
		path = *subcommand
	}
	if subcommandGroup != nil {
		path += "/" + *subcommandGroup
	}
	return path
}

func (m *CommandMap) AddAll(c []Command) {
	for _, cmd := range c {
		m.commands[cmd.Create.Name()] = cmd
	}
}
