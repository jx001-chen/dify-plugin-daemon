package backwards_invocation

import (
	"testing"

	"github.com/langgenius/dify-plugin-daemon/internal/core/dify_invocation"
	"github.com/langgenius/dify-plugin-daemon/internal/core/plugin_manager/positive_manager"
	"github.com/langgenius/dify-plugin-daemon/internal/types/entities/plugin_entities"
)

type TPluginRuntime struct {
	plugin_entities.PluginRuntime
	positive_manager.PositivePluginRuntime
}

func (r *TPluginRuntime) InitEnvironment() error {
	return nil
}

func (r *TPluginRuntime) Checksum() string {
	return ""
}

func (r *TPluginRuntime) Identity() (string, error) {
	return "", nil
}

func (r *TPluginRuntime) StartPlugin() error {
	return nil
}

func (r *TPluginRuntime) Type() plugin_entities.PluginRuntimeType {
	return plugin_entities.PLUGIN_RUNTIME_TYPE_LOCAL
}

func (r *TPluginRuntime) Wait() (<-chan bool, error) {
	return nil, nil
}

func TestBackwardsInvocationAllPermittedPermission(t *testing.T) {
	all_permitted_runtime := plugin_entities.PluginDeclaration{
		PluginDeclarationWithoutAdvancedFields: plugin_entities.PluginDeclarationWithoutAdvancedFields{
			Resource: plugin_entities.PluginResourceRequirement{
				Permission: &plugin_entities.PluginPermissionRequirement{
					Tool: &plugin_entities.PluginPermissionToolRequirement{
						Enabled: true,
					},
					Model: &plugin_entities.PluginPermissionModelRequirement{
						Enabled:       true,
						LLM:           true,
						TextEmbedding: true,
						Rerank:        true,
						Moderation:    true,
						TTS:           true,
						Speech2text:   true,
					},
					Node: &plugin_entities.PluginPermissionNodeRequirement{
						Enabled: true,
					},
					App: &plugin_entities.PluginPermissionAppRequirement{
						Enabled: true,
					},
				},
			},
		},
	}

	invoke_llm_request := NewBackwardsInvocation(dify_invocation.INVOKE_TYPE_LLM, "", nil, nil, nil)
	if err := checkPermission(&all_permitted_runtime, invoke_llm_request); err != nil {
		t.Errorf("checkPermission failed: %s", err.Error())
	}

	invoke_text_embedding_request := NewBackwardsInvocation(dify_invocation.INVOKE_TYPE_TEXT_EMBEDDING, "", nil, nil, nil)
	if err := checkPermission(&all_permitted_runtime, invoke_text_embedding_request); err != nil {
		t.Errorf("checkPermission failed: %s", err.Error())
	}

	invoke_rerank_request := NewBackwardsInvocation(dify_invocation.INVOKE_TYPE_RERANK, "", nil, nil, nil)
	if err := checkPermission(&all_permitted_runtime, invoke_rerank_request); err != nil {
		t.Errorf("checkPermission failed: %s", err.Error())
	}

	invoke_tts_request := NewBackwardsInvocation(dify_invocation.INVOKE_TYPE_TTS, "", nil, nil, nil)
	if err := checkPermission(&all_permitted_runtime, invoke_tts_request); err != nil {
		t.Errorf("checkPermission failed: %s", err.Error())
	}

	invoke_speech2text_request := NewBackwardsInvocation(dify_invocation.INVOKE_TYPE_SPEECH2TEXT, "", nil, nil, nil)
	if err := checkPermission(&all_permitted_runtime, invoke_speech2text_request); err != nil {
		t.Errorf("checkPermission failed: %s", err.Error())
	}

	invoke_moderation_request := NewBackwardsInvocation(dify_invocation.INVOKE_TYPE_MODERATION, "", nil, nil, nil)
	if err := checkPermission(&all_permitted_runtime, invoke_moderation_request); err != nil {
		t.Errorf("checkPermission failed: %s", err.Error())
	}

	invoke_tool_request := NewBackwardsInvocation(dify_invocation.INVOKE_TYPE_TOOL, "", nil, nil, nil)
	if err := checkPermission(&all_permitted_runtime, invoke_tool_request); err != nil {
		t.Errorf("checkPermission failed: %s", err.Error())
	}

	invoke_node_request := NewBackwardsInvocation(dify_invocation.INVOKE_TYPE_NODE, "", nil, nil, nil)
	if err := checkPermission(&all_permitted_runtime, invoke_node_request); err != nil {
		t.Errorf("checkPermission failed: %s", err.Error())
	}

	invoke_app_request := NewBackwardsInvocation(dify_invocation.INVOKE_TYPE_APP, "", nil, nil, nil)
	if err := checkPermission(&all_permitted_runtime, invoke_app_request); err != nil {
		t.Errorf("checkPermission failed: %s", err.Error())
	}
}

func TestBackwardsInvocationAllDeniedPermission(t *testing.T) {
	all_denied_runtime := plugin_entities.PluginDeclaration{
		PluginDeclarationWithoutAdvancedFields: plugin_entities.PluginDeclarationWithoutAdvancedFields{
			Resource: plugin_entities.PluginResourceRequirement{},
		},
	}

	invoke_llm_request := NewBackwardsInvocation(dify_invocation.INVOKE_TYPE_LLM, "", nil, nil, nil)
	if err := checkPermission(&all_denied_runtime, invoke_llm_request); err == nil {
		t.Errorf("checkPermission failed: expected error, got nil")
	}

	invoke_text_embedding_request := NewBackwardsInvocation(dify_invocation.INVOKE_TYPE_TEXT_EMBEDDING, "", nil, nil, nil)
	if err := checkPermission(&all_denied_runtime, invoke_text_embedding_request); err == nil {
		t.Errorf("checkPermission failed: expected error, got nil")
	}

	invoke_rerank_request := NewBackwardsInvocation(dify_invocation.INVOKE_TYPE_RERANK, "", nil, nil, nil)
	if err := checkPermission(&all_denied_runtime, invoke_rerank_request); err == nil {
		t.Errorf("checkPermission failed: expected error, got nil")
	}

	invoke_tts_request := NewBackwardsInvocation(dify_invocation.INVOKE_TYPE_TTS, "", nil, nil, nil)
	if err := checkPermission(&all_denied_runtime, invoke_tts_request); err == nil {
		t.Errorf("checkPermission failed: expected error, got nil")
	}

	invoke_speech2text_request := NewBackwardsInvocation(dify_invocation.INVOKE_TYPE_SPEECH2TEXT, "", nil, nil, nil)
	if err := checkPermission(&all_denied_runtime, invoke_speech2text_request); err == nil {
		t.Errorf("checkPermission failed: expected error, got nil")
	}

	invoke_moderation_request := NewBackwardsInvocation(dify_invocation.INVOKE_TYPE_MODERATION, "", nil, nil, nil)
	if err := checkPermission(&all_denied_runtime, invoke_moderation_request); err == nil {
		t.Errorf("checkPermission failed: expected error, got nil")
	}

	invoke_tool_request := NewBackwardsInvocation(dify_invocation.INVOKE_TYPE_TOOL, "", nil, nil, nil)
	if err := checkPermission(&all_denied_runtime, invoke_tool_request); err == nil {
		t.Errorf("checkPermission failed: expected error, got nil")
	}

	invoke_node_request := NewBackwardsInvocation(dify_invocation.INVOKE_TYPE_NODE, "", nil, nil, nil)
	if err := checkPermission(&all_denied_runtime, invoke_node_request); err == nil {
		t.Errorf("checkPermission failed: expected error, got nil")
	}

	invoke_app_request := NewBackwardsInvocation(dify_invocation.INVOKE_TYPE_APP, "", nil, nil, nil)
	if err := checkPermission(&all_denied_runtime, invoke_app_request); err == nil {
		t.Errorf("checkPermission failed: expected error, got nil")
	}
}
