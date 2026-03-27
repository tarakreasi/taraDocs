<script setup>
import { computed } from 'vue';
import { marked } from 'marked';
import { router } from '@inertiajs/vue3';

const props = defineProps({
    content: { type: String, default: '' },
    displayName: { type: String, default: '' },
    isEditing: { type: Boolean, default: false },
    extension: { type: String, default: '.md' },
});

const emit = defineEmits(['edit']);

const renderedHtml = computed(() => {
    if (!props.content) return '<p class="text-slate-400 italic">This page is empty.</p>';
    if (props.extension === '.html') return props.content; // Render raw HTML
    return marked.parse(props.content); // Render Markdown
});

// Handle clicks on internal markdown links
const handleContentClick = (event) => {
    const link = event.target.closest('a');
    if (!link) return;

    const href = link.getAttribute('href');
    if (!href) return;

    try {
        const url = new URL(href, window.location.href);
        if (url.origin === window.location.origin) {
            event.preventDefault();
            router.visit(url.href);
        }
        // external links open normally (no preventDefault)
    } catch {
        // relative path - let browser handle
    }
};
</script>

<template>
    <main class="flex-1 flex flex-col bg-slate-50 dark:bg-[#0F172A] overflow-hidden">

        <!-- Toolbar -->
        <header v-if="displayName" class="h-14 flex items-center justify-between px-8 shrink-0 border-b border-slate-200 dark:border-white/5">
            <div class="flex items-center gap-2 text-xs text-slate-400">
                <span class="material-symbols-outlined text-[16px]">description</span>
                <span class="font-medium text-slate-600 dark:text-slate-300">{{ displayName }}</span>
            </div>

            <button
                v-if="!isEditing"
                @click="emit('edit')"
                class="flex items-center gap-1.5 h-8 px-3 rounded-lg text-[12px] font-semibold bg-indigo-500 text-white hover:bg-indigo-600 transition-colors shadow-sm"
            >
                <span class="material-symbols-outlined text-[16px]">edit</span>
                Edit
            </button>
        </header>

        <!-- Empty state when no doc is selected -->
        <div v-if="!displayName" class="flex-1 flex flex-col items-center justify-center opacity-40">
            <div class="size-16 rounded-2xl bg-slate-100 dark:bg-white/5 flex items-center justify-center mb-4">
                <span class="material-symbols-outlined text-3xl">article</span>
            </div>
            <h2 class="text-sm font-medium text-slate-800 dark:text-white mb-1">Select a document</h2>
            <p class="text-xs text-slate-400">Choose from the list on the left</p>
        </div>

        <!-- Content -->
        <div v-else class="flex-1 overflow-y-auto">
            <article
                class="max-w-4xl mx-auto px-10 py-12 prose prose-slate dark:prose-invert prose-headings:font-sans prose-headings:font-bold prose-a:text-indigo-500 hover:prose-a:text-indigo-600 prose-code:bg-slate-100 dark:prose-code:bg-white/10 prose-code:px-1 prose-code:py-0.5 prose-code:rounded prose-code:before:content-none prose-code:after:content-none [&_pre_code]:bg-transparent [&_pre_code]:p-0 max-w-none"
                v-html="renderedHtml"
                @click="handleContentClick"
            ></article>
        </div>
    </main>
</template>
