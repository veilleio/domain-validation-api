#!/usr/bin/env bash
# Domain Validation API – cURL example
# Usage: VEILLE_API_KEY=your-key bash intelligence_domain.sh
set -euo pipefail

curl -sS -G "https://api.veille.io/v1/intelligence/domain" \
  -H "x-api-key: ${VEILLE_API_KEY}" \
  --data-urlencode query=YOUR_VALUE
