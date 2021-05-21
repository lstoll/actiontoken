

curl --request GET \
    --url https://api.github.com/repos/${GITHUB_REPOSITORY} \
    --header "authorization: Bearer $(cat ~/.github_token)" \
    --header 'content-type: application/json'

curl -X GET \
    -H "Authorization: Bearer $(cat ~/.github_token)" \
    -H 'Accept: application/vnd.github.v3.full+json' \
    https://api.github.com/

curl -X GET \
    -H "Authorization: Bearer $(cat ~/.github_token)" \
    -H 'Accept: application/vnd.github.v3.full+json' \
    "https://api.github.com/repos/${GITHUB_REPOSITORY}/actions"

curl -X POST \
    -H "Authorization: Bearer $(cat ~/.github_token)" \
    -H 'Accept: application/vnd.github.v3.full+json' \
    "https://api.github.com/repos/${GITHUB_REPOSITORY}/check-runs" \
    -d "{\"name\":\"ghactionsid\",\"head_sha\":\"$GITHUB_SHA\"}"

# https://stackoverflow.com/questions/9765453/is-gits-semi-secret-empty-tree-object-reliable-and-why-is-there-not-a-symbolic
curl -X POST \
    -H "Authorization: Bearer $(cat ~/.github_token)" \
    -H 'Accept: application/vnd.github.v3.full+json' \
    "https://api.github.com/repos/${GITHUB_REPOSITORY}/git/commits" \
    -d '{"message": "verifying token write access", "tree": "4b825dc642cb6eb9a060e54bf8d69288fbee4904", "parents": []}'
