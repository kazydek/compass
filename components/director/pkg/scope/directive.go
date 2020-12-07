package scope

import (
	"context"
	"strings"

	"github.com/kyma-incubator/compass/components/director/pkg/apperrors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/pkg/errors"
)

//go:generate mockery -name=ScopesGetter -output=automock -outpkg=automock -case=underscore
type ScopesGetter interface {
	GetRequiredScopes(scopesDefinition string) ([]string, error)
}

type directive struct {
	scopesGetter ScopesGetter
	prefix       string
}

func NewDirective(getter ScopesGetter, prefix string) *directive {
	return &directive{
		scopesGetter: getter,
		prefix:       prefix,
	}
}

func (d *directive) VerifyScopes(ctx context.Context, obj interface{}, next graphql.Resolver, scopesDefinition string) (interface{}, error) {
	actualScopes, err := LoadFromContext(ctx)
	if err != nil {
		return nil, err
	}
	requiredScopes, err := d.scopesGetter.GetRequiredScopes(scopesDefinition)
	if err != nil {
		return nil, errors.Wrap(err, "while getting required scopes")
	}

	if !d.matches(actualScopes, requiredScopes) {
		return nil, apperrors.NewInsufficientScopesError(requiredScopes, actualScopes)
	}
	return next(ctx)
}

func (d *directive) matches(actual []string, required []string) bool {
	actMap := make(map[string]interface{})

	for _, a := range actual {
		a = strings.TrimPrefix(a, d.prefix)
		actMap[a] = struct{}{}
	}
	for _, r := range required {
		_, ex := actMap[r]
		if !ex {
			return false
		}
	}
	return true
}
