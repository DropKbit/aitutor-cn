package lesson_test

import (
	"testing"
	"unicode"

	_ "github.com/DropKbit/aitutor-cn/internal/content/advanced"
	_ "github.com/DropKbit/aitutor-cn/internal/content/beginner"
	_ "github.com/DropKbit/aitutor-cn/internal/content/intermediate"
	"github.com/DropKbit/aitutor-cn/internal/lesson"
	"github.com/DropKbit/aitutor-cn/internal/ui"
	"github.com/DropKbit/aitutor-cn/pkg/types"
)

func TestTierStringsAreLocalized(t *testing.T) {
	if got := types.Beginner.String(); got != "初级" {
		t.Fatalf("Beginner.String() = %q, want %q", got, "初级")
	}
	if got := types.Intermediate.String(); got != "中级" {
		t.Fatalf("Intermediate.String() = %q, want %q", got, "中级")
	}
	if got := types.Advanced.String(); got != "高级" {
		t.Fatalf("Advanced.String() = %q, want %q", got, "高级")
	}
}

func TestPhaseStringsAreLocalized(t *testing.T) {
	if got := lesson.PhaseTheory.String(); got != "理论" {
		t.Fatalf("PhaseTheory.String() = %q, want %q", got, "理论")
	}
	if got := lesson.PhaseViz.String(); got != "可视化" {
		t.Fatalf("PhaseViz.String() = %q, want %q", got, "可视化")
	}
	if got := lesson.PhaseQuiz.String(); got != "测验" {
		t.Fatalf("PhaseQuiz.String() = %q, want %q", got, "测验")
	}
	if got := lesson.PhaseComplete.String(); got != "完成" {
		t.Fatalf("PhaseComplete.String() = %q, want %q", got, "完成")
	}
}

func TestAllLessonsRemainRegisteredAndLocalized(t *testing.T) {
	lessons := lesson.All()
	if len(lessons) != 17 {
		t.Fatalf("lesson count = %d, want 17", len(lessons))
	}

	for _, def := range lessons {
		if !containsHan(def.Title) {
			t.Fatalf("lesson %d title is not localized: %q", def.ID, def.Title)
		}
		if !containsHan(def.Summary) {
			t.Fatalf("lesson %d summary is not localized: %q", def.ID, def.Summary)
		}
		if len(def.Theory) == 0 {
			t.Fatalf("lesson %d has no theory blocks", def.ID)
		}
		if len(def.Questions) == 0 {
			t.Fatalf("lesson %d has no quiz questions", def.ID)
		}
	}
}

func TestSidebarViewHandlesChineseTitles(t *testing.T) {
	sidebar := ui.SidebarModel{
		Width: 10,
		Lessons: []types.LessonDef{
			{ID: 1, Title: "这是一个很长的中文课程标题", Tier: types.Beginner},
		},
		Completed: make(map[int]bool),
	}

	_ = sidebar.View()
}

func containsHan(s string) bool {
	for _, r := range s {
		if unicode.Is(unicode.Han, r) {
			return true
		}
	}
	return false
}
