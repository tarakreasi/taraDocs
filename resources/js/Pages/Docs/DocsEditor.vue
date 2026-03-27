<script setup>
import { ref, watch, onBeforeUnmount } from 'vue';
import NoteEditor from '@/Components/NoteEditor.vue';
import TurndownService from 'turndown';
import axios from 'axios';

const props = defineProps({
    currentPath: { type: String, required: true },
    // The raw content from the file
    content: { type: String, default: '' },
    extension: { type: String, default: '.md' },
});

const emit = defineEmits(['cancel', 'saved']);

// Convert incoming raw Markdown to HTML for Tiptap on entry
import { marked } from 'marked';

// editorHtml holds the Tiptap rich-text value (HTML format internally)
const editorHtml = ref('');
const isSaving = ref(false);
const lastSavedAt = ref(null);
const saveStatus = ref('');

// When the editor isを開く opened, load content
watch(() => props.content, (content) => {
    if (content) {
        if (props.extension === '.html') {
            editorHtml.value = content; // Load pure HTML
        } else {
            editorHtml.value = marked.parse(content); // Markdown -> HTML
        }
    }
}, { immediate: true });

// Auto-save debounce
let autoSaveTimer = null;
const onEditorChange = () => {
    clearTimeout(autoSaveTimer);
    saveStatus.value = 'Unsaved';
    autoSaveTimer = setTimeout(() => saveDoc(true), 1500);
};

const saveDoc = async (isAutoSave = false) => {
    if (!props.currentPath) return;
    isSaving.value = true;
    saveStatus.value = 'Saving...';

    try {
        let finalContent = editorHtml.value;
        
        // If it's a markdown file, convert Tiptap HTML → Markdown
        if (props.extension === '.md') {
            const td = new TurndownService({ headingStyle: 'atx', codeBlockStyle: 'fenced' });
            finalContent = td.turndown(editorHtml.value);
        }

        await axios.post('/docs/save', {
            path: props.currentPath,
            content: finalContent,
            extension: props.extension,
        });

        lastSavedAt.value = new Date();
        saveStatus.value = 'Saved';
        emit('saved', finalContent);
    } catch (err) {
        console.error(err);
        saveStatus.value = 'Error saving';
        if (!isAutoSave) alert('Failed to save document.');
    } finally {
        isSaving.value = false;
    }
};

// Keyboard shortcut: Ctrl+S to save, Escape to cancel
const handleKeydown = (e) => {
    if ((e.ctrlKey || e.metaKey) && e.key === 's') {
        e.preventDefault();
        saveDoc(false);
    }
    if (e.key === 'Escape') {
        emit('cancel');
    }
};

import { onMounted } from 'vue';
onMounted(() => window.addEventListener('keydown', handleKeydown));
onBeforeUnmount(() => {
    window.removeEventListener('keydown', handleKeydown);
    clearTimeout(autoSaveTimer);
});
</script>

<template>
    <main class="flex-1 flex flex-col bg-slate-50 dark:bg-[#0F172A] overflow-hidden">

        <!-- Editor Toolbar -->
        <header class="h-14 flex items-center justify-between px-6 shrink-0 border-b border-slate-200 dark:border-white/5">
            <div class="flex items-center gap-2 text-xs font-mono text-slate-400">
                <span class="material-symbols-outlined text-[16px] text-indigo-500">edit</span>
                <span class="text-slate-600 dark:text-slate-300">Editing: <strong>{{ currentPath }}{{ extension }}</strong></span>
            </div>

            <div class="flex items-center gap-2">
                <!-- Save Status -->
                <span class="text-[11px] font-mono mr-2"
                    :class="{
                        'text-blue-500': saveStatus === 'Saving...',
                        'text-green-500': saveStatus === 'Saved',
                        'text-orange-400': saveStatus === 'Unsaved',
                        'text-red-500': saveStatus.startsWith('Error'),
                    }"
                >{{ saveStatus }}</span>

                <!-- Save -->
                <button
                    @click="saveDoc(false)"
                    :disabled="isSaving"
                    class="flex items-center gap-1.5 h-8 px-3 rounded-lg text-[12px] font-semibold bg-indigo-500 text-white hover:bg-indigo-600 disabled:opacity-50 transition-colors shadow-sm"
                >
                    <span class="material-symbols-outlined text-[15px]">save</span>
                    Save
                </button>

                <!-- Cancel -->
                <button
                    @click="emit('cancel')"
                    class="flex items-center gap-1.5 h-8 px-3 rounded-lg text-[12px] font-semibold bg-slate-100 dark:bg-white/10 text-slate-600 dark:text-slate-300 hover:bg-slate-200 dark:hover:bg-white/20 transition-colors"
                >
                    Cancel
                </button>
            </div>
        </header>

        <!-- Tiptap Editor -->
        <div class="flex-1 overflow-hidden flex flex-col">
            <NoteEditor
                v-model="editorHtml"
                :save-status="saveStatus"
                :last-saved-at="lastSavedAt"
                @save="onEditorChange"
            />
        </div>
    </main>
</template>
