# JUXON-AI ad-free build

This fork packages the CLIProxyAPI v8.0.44 backend with an ad-free management panel based on v1.27.2. Provider configuration, both regional Kimi OAuth logins, plugin support, and proxy API behavior are retained. Commercial campaign banners, registration referrals, proxy recommendations, sponsored OAuth presentation, and README promotions are removed. Upstream copyright and licenses are retained.

The panel source is pinned to `JUXON-AI/Cli-Proxy-API-Management-Center@cef87e5d272322cd9c47280ebf042a52d2198f9a`. `Dockerfile.adfree` also verifies the source archive digest, runs the panel tests/lint/build, and includes both licenses in `/CLIProxyAPI/licenses`.

## Build and publish

Run on an authorized Linux/amd64 build server using a clean checkout of the intended backend commit:

```sh
SOURCE_SHA=$(git rev-parse HEAD)
SHORT_SHA=$(git rev-parse --short=12 HEAD)
IMAGE=registry.cn-hangzhou.aliyuncs.com/skfskf/claude-proxy:8.0.44-adfree-${SHORT_SHA}
docker buildx build --platform linux/amd64 --provenance=false --load \
  --file Dockerfile.adfree \
  --build-arg COMMIT="$SOURCE_SHA" \
  --build-arg BUILD_DATE="$(date -u +%Y-%m-%dT%H:%M:%SZ)" \
  --tag "$IMAGE" .
docker push "$IMAGE"
docker buildx imagetools inspect "$IMAGE"
```

Use the immutable tag and registry digest for any later deployment. Publishing does not change the running VPS service.

## Runtime configuration and upgrades

The panel is already present at `/CLIProxyAPI/static/management.html`; first access does not require a GitHub download. The example config disables periodic panel updates. Default release and fallback download sources point to the JUXON-AI panel fork.

When reusing an existing mounted config, set these fields explicitly before deployment:

```yaml
management:
  disable-control-panel: false
  disable-auto-update-panel: true
  panel-github-repository: "https://github.com/JUXON-AI/Cli-Proxy-API-Management-Center"
```

Keep the existing management secret and client API keys in their private runtime configuration. The image contains no account credentials or keys. An explicit runtime override to another panel repository can download that repository's panel; review that choice before enabling updates. Existing `/quick-start` bookmarks redirect to `/ai-providers`, where all provider groups remain configurable.

No model inference is needed to validate panel delivery or API authentication. A browser interaction check is a separate acceptance step from the server-side render, build, and HTTP checks.
