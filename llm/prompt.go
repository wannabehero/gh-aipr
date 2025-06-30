package llm

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

const SYSTEM_PROMPT = `You are a tool that summarises git commit messages
and git diffs to generate a descriptive Pull Request title and body.

Title describes the changes in no more than 10 words.
Start with a verb like "create", "fix", "update", etc.
Put a random emoji that describes the context in front of the title.
For instance: '👾 Update the build CI workflow'

Your responses are concise and to the point`

const BODY_TEMPLATE_PROMPT = `IMPORTANT FOR THE BODY: Format your response according to the provided PR template structure.
Fill in the sections appropriately while maintaining the template format.

PR Template:
%s`

const USER_PROMPT = `
Using the following commit messages and a diff as a context
generate a descriptive concise Pull Request title and body that summarizes the changes.

For the title include an emoji that describes the context to the start.

For the body include a brief summary section with 1-3 bullet points describing the key changes.
Format it in Markdown with proper headings.

%s

Commit messages:
%s

Diff:
%s
`

func getUserPrompt(commits []string, diff string, prTemplate string) string {
	var prTemplatePrompt string
	if prTemplate != "" {
		prTemplatePrompt = fmt.Sprintf(BODY_TEMPLATE_PROMPT, prTemplate)
	} else {
		prTemplatePrompt = ""
	}

	prompt := fmt.Sprintf(USER_PROMPT, prTemplatePrompt, strings.Join(commits, "\n"), diff)
	return prompt
}

func getSystemPrompt() string {
	basePrompt := SYSTEM_PROMPT

	if override := viper.GetString("system_prompt_override"); override != "" {
		basePrompt = override
	}

	if instructions := viper.GetString("additional_instructions"); instructions != "" {
		return basePrompt + "\n\n## Additional User-Customized Instructions\n" + instructions
	}

	return basePrompt
}

type Response struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func getProvider(ctx context.Context) LLMProvider {
	if key := os.Getenv("GEMINI_API_KEY"); key != "" {
		provider := NewGeminiProvider(key, ctx)
		if provider != nil {
			return provider
		}
	}

	if key := os.Getenv("OPENAI_API_KEY"); key != "" {
		provider := NewOpenaiProvider(key)
		if provider != nil {
			return provider
		}
	}

	if key := os.Getenv("ANTHROPIC_API_KEY"); key != "" {
		provider := NewAnthropicProvider(key)
		if provider != nil {
			return provider
		}
	}

	return nil
}

func GenerateTitleAndBody(commits []string, diff string, template string, ctx context.Context) (*string, *string) {
	provider := getProvider(ctx)
	if provider == nil {
		return nil, nil
	}
	return provider.GenerateTitleAndBody(commits, diff, template, ctx)
}
