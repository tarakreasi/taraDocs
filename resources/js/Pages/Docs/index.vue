<script setup>
import { ref, onMounted, watch } from 'vue';
import { Head, router } from '@inertiajs/vue3';
import axios from 'axios';
import DocsSidebar from './DocsSidebar.vue';
import DocsFileList from './DocsFileList.vue';
import DocsReader from './DocsReader.vue';
import DocsEditor from './DocsEditor.vue';

// ──────────────────────────────────────────────
// Props from Go backend (via Inertia)
// ──────────────────────────────────────────────
const props = defineProps({
    content: { type: String, default: '' },
    currentPath: { type: String, default: '' },
    displayName: { type: String, default: '' },
    extension: { type: String, default: '.md' },
});

// ──────────────────────────────────────────────
// Navigation: dynamic docs list from API
// ──────────────────────────────────────────────
const allDocs = ref([]);
const isNavLoading = ref(true);

const fetchNavigation = async () => {
    try {
        const res = await axios.get('/docs/api/navigation');
        if (res.data?.success) {
            allDocs.value = res.data.data;
            syncCategoryFromPath(); // Sync category automatically upon data load
        }
    } catch (e) {
        console.error('Failed to load navigation', e);
    } finally {
        isNavLoading.value = false;
    }
};

// ──────────────────────────────────────────────
// Sidebar: selected category
// ──────────────────────────────────────────────
const selectedCategory = ref(null);

// Automatically set category based on the open file path
const syncCategoryFromPath = () => {
    if (!props.currentPath || allDocs.value.length === 0) return;
    const currentDoc = allDocs.value.find(d => d.path === props.currentPath || d.name === props.currentPath);
    if (currentDoc) {
        selectedCategory.value = currentDoc.category;
    }
};

// Re-sync if the user navigates to a new route in the app shell
watch(() => props.currentPath, () => {
    syncCategoryFromPath();
});

const onSelectCategory = (category) => {
    selectedCategory.value = category;
};

// ──────────────────────────────────────────────
// Editor state
// ──────────────────────────────────────────────
const isEditing = ref(false);

const onEdit = () => {
    isEditing.value = true;
};

const onCancel = () => {
    isEditing.value = false;
};

// When doc is saved, reload content from server, exit edit mode
const onSaved = () => {
    isEditing.value = false;
    router.reload({ only: ['content'], preserveScroll: true, preserveState: true });
};

// ──────────────────────────────────────────────
// Lifecycle
// ──────────────────────────────────────────────
onMounted(() => {
    fetchNavigation();
});

// Noise style (subtle texture)
const noiseStyle = {
    backgroundImage: `url("data:image/svg+xml,%3Csvg viewBox='0 0 200 200' xmlns='http://www.w3.org/2000/svg'%3E%3Cfilter id='n'%3E%3CfeTurbulence type='fractalNoise' baseFrequency='0.65' numOctaves='3' stitchTiles='stitch'/%3E%3C/filter%3E%3Crect width='100%25' height='100%25' filter='url(%23n)' opacity='0.5'/%3E%3C/svg%3E")`,
};
</script>

<template>
    <div>
        <Head :title="`${displayName || 'Docs'} — TaraDocs`" />

        <div class="h-screen flex font-sans text-slate-800 dark:text-white bg-slate-50 dark:bg-[#0F172A] transition-colors overflow-hidden relative">

            <!-- Noise & Ambient background -->
            <div class="fixed inset-0 z-0 opacity-[0.03] dark:opacity-[0.02] pointer-events-none" :style="noiseStyle" />
            <div class="fixed inset-0 pointer-events-none z-0"
                 style="background: radial-gradient(circle at 50% 0%, rgba(99,102,241,0.04) 0%, transparent 60%);">
                <div class="absolute -top-1/4 -left-1/4 w-1/2 h-1/2 bg-indigo-500/5 blur-[120px] rounded-full" />
                <div class="absolute -bottom-1/4 -right-1/4 w-1/2 h-1/2 bg-rose-500/5 blur-[120px] rounded-full" />
            </div>

            <!-- App Shell: 3-column layout -->
            <div class="relative z-10 flex w-full h-full overflow-hidden">

                <!-- 1. Left: Category Sidebar -->
                <DocsSidebar
                    :all-docs="allDocs"
                    :selected-category="selectedCategory"
                    @select-category="onSelectCategory"
                    @category-created="fetchNavigation"
                />

                <!-- 2. Middle: File List -->
                <DocsFileList
                    :all-docs="allDocs"
                    :selected-category="selectedCategory"
                    :current-path="currentPath"
                    @file-created="fetchNavigation"
                />

                <!-- 3. Right: Reader or Editor -->
                <DocsReader
                    v-if="!isEditing"
                    :content="content"
                    :display-name="displayName"
                    :is-editing="isEditing"
                    :extension="extension"
                    @edit="onEdit"
                />

                <DocsEditor
                    v-else
                    :content="content"
                    :current-path="currentPath"
                    :extension="extension"
                    @cancel="onCancel"
                    @saved="onSaved"
                />

            </div>
        </div>
    </div>
</template>
