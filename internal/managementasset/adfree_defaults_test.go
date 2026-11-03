package managementasset

import (
	"testing"

	"github.com/router-for-me/CLIProxyAPI/v8/internal/config"
)

func TestAdFreePanelSources(t *testing.T) {
	const release = "https://api.github.com/repos/JUXON-AI/Cli-Proxy-API-Management-Center/releases/latest"
	if got := resolveReleaseURL(config.DefaultPanelGitHubRepository); got != release {
		t.Fatalf("default repository resolves to %q, want organization releases %q", got, release)
	}
	if got := resolveReleaseURL(""); got != release {
		t.Fatalf("empty repository resolves to %q, want %q", got, release)
	}
	const fallback = "https://github.com/JUXON-AI/Cli-Proxy-API-Management-Center/releases/latest/download/management.html"
	if defaultManagementFallbackURL != fallback {
		t.Fatalf("fallback = %q, want organization panel %q", defaultManagementFallbackURL, fallback)
	}
}
