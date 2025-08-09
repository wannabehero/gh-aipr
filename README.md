# gh-aipr

GitHub CLI extension to generate pull request titles and descriptions
based on your current branch commits and the diff.

## Installation

### GitHub CLI Extension

Make sure the CLI is installed: https://cli.github.com/

```
gh extension install wannabehero/gh-aipr
```

### Nix (Home Manager)

If using flakes, enable the overlay so `pkgs.gh-aipr` is available:

```nix
# flake inputs
inputs.gh-aipr.url = "github:wannabehero/gh-aipr";

# when constructing pkgs
pkgs = import nixpkgs {
  inherit system;
  overlays = [ gh-aipr.overlays.pkgs ];
};
```

Then in your Home Manager configuration:

```nix
programs.gh = {
  enable = true;
  extensions = [ pkgs.gh-aipr ];
};
```

Or install directly:

```bash
# Build and run once
nix run github:wannabehero/gh-aipr

# Install to profile
nix profile install github:wannabehero/gh-aipr
```

## Configuration

To make the best use of the tool set one of the following:
- `OPENAI_API_KEY`
- `ANTHROPIC_API_KEY`
- `GEMINI_API_KEY`

in your environment variables so it can use the LLM
to generate relevant title automatically
based on your current branch commits.

The tool uses some sensible defaults for the models
but you can override them in the config file.

See more in [docs/config.md](docs/config.md).

## Usage

```
gh aipr <other args that gh pr create supports>
```

This will follow you through interactive create process
and your final PR title and body will be automatically generated and set.
