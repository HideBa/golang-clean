package domain

import (
	"reflect"
	"testing"
)

func TestNewArticleModel(t *testing.T) {
	type args struct {
		content string
		userID  uint64
	}
	tests := []struct {
		name string
		args args
		want *ArticleModel
	}{
		{name: "it should generate new article model with content and userID",
			args: args{content: "hello world", userID: 1},
			want: &ArticleModel{Content: "hello world", UserID: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewArticleModel(tt.args.content, tt.args.userID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewArticleModel() = %v, want %v", got, tt.want)
			}
		})
	}
}
