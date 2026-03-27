<script setup>
import { computed, ref } from 'vue';
import { Link } from '@inertiajs/vue3';
import axios from 'axios';

const props = defineProps({
    allDocs: { type: Array, default: () => [] },
    selectedCategory: { type: String, default: null },
    currentPath: { type: String, default: '' },
});

const searchQuery = ref('');

const categoryLabel = computed(() => props.selectedCategory || 'All Documentation');

const filteredDocs = computed(() => {
    let list = props.allDocs;

    if (props.selectedCategory) {
        list = list.filter(doc => doc.category === props.selectedCategory);
    }

    if (searchQuery.value.trim()) {
        const q = searchQuery.value.toLowerCase();
        list = list.filter(doc =>
            doc.title.toLowerCase().includes(q) ||
            doc.path.toLowerCase().includes(q)
        );
    }

    return list;
});

const docHref = (doc) => `/docs/${doc.path}`;

const emit = defineEmits(['file-created']);
const isCreatingFile = ref(false);

const createNewFile = async () => {
    const fileName = prompt(`Enter new file name in folder "${props.selectedCategory || 'Uncategorized'}":\n(Add .html if you want an HTML file, defaults to .md)`);
    if (!fileName) return;

    let folderPath = "";
    if (props.selectedCategory && props.selectedCategory !== "Uncategorized") {
        folderPath = props.selectedCategory.replace(/ \/ /g, "/"); // Convert "A / B" to "A/B"
    }

    try {
        isCreatingFile.value = true;
        const response = await axios.post('/docs/api/create-file', {
            folderPath: folderPath,
            fileName: fileName,
        });

        if (response.data.success) {
            emit('file-created'); // Tell parent to reload 
        }
    } catch (error) {
        alert(error.response?.data?.error || "Failed to create file");
    } finally {
        isCreatingFile.value = false;
    }
};
</script>

<template>
    <div class="w-[300px] flex flex-col bg-white/50 dark:bg-[#0b1121]/80 border-r border-slate-200 dark:border-white/5 shrink-0">

        <!-- Header -->
        <div class="h-16 flex items-center justify-between px-5 shrink-0 border-b border-slate-200 dark:border-white/5">
            <h2 class="font-sans font-bold text-[13px] uppercase tracking-wide text-slate-700 dark:text-slate-300 truncate">
                {{ categoryLabel }}
            </h2>
            <button 
                @click="createNewFile"
                :disabled="isCreatingFile || !props.selectedCategory || props.selectedCategory === 'Uncategorized'"
                class="size-6 rounded flex items-center justify-center text-slate-400 hover:text-indigo-500 hover:bg-slate-200 dark:hover:bg-white/10 transition-colors disabled:opacity-30 disabled:hover:bg-transparent"
                title="New Markdown Document"
            >
                <span class="material-symbols-outlined text-[16px]">note_add</span>
            </button>
        </div>

        <!-- Search -->
        <div class="px-4 py-3 shrink-0">
            <div class="relative">
                <span class="absolute left-3 top-1/2 -translate-y-1/2 material-symbols-outlined text-slate-400 text-[18px]">search</span>
                <input
                    v-model="searchQuery"
                    type="text"
                    placeholder="Search docs..."
                    class="w-full h-9 bg-slate-100 dark:bg-white/5 border-none rounded-lg pl-9 pr-3 text-sm text-slate-700 dark:text-slate-200 placeholder:text-slate-400 focus:ring-1 focus:ring-indigo-500 transition-all outline-none"
                />
            </div>
        </div>

        <!-- Doc List -->
        <div class="flex-1 overflow-y-auto px-3 pb-4">

            <!-- Empty state -->
            <div v-if="filteredDocs.length === 0" class="flex flex-col items-center justify-center h-full text-center opacity-50">
                <span class="material-symbols-outlined text-4xl mb-2 text-slate-300">description</span>
                <p class="text-xs text-slate-500">No docs found</p>
            </div>

            <!-- Cards -->
            <div v-else class="space-y-1 pt-1">
                <Link
                    v-for="doc in filteredDocs"
                    :key="doc.path"
                    preserve-state
                    preserve-scroll
                    :href="docHref(doc)"
                    class="block p-3 rounded-lg cursor-pointer transition-all duration-200 group"
                    :class="currentPath === doc.path
                        ? 'bg-white dark:bg-white/10 shadow-md ring-1 ring-black/5 dark:ring-white/5'
                        : 'hover:bg-white/60 dark:hover:bg-white/5 hover:shadow-sm'"
                >
                    <h3 class="font-medium text-[13px] tracking-tight line-clamp-2 leading-snug"
                        :class="currentPath === doc.path ? 'text-indigo-600 dark:text-indigo-400' : 'text-slate-800 dark:text-white'"
                    >
                        {{ doc.title }}
                        <span v-if="doc.extension === '.html'" class="inline-block ml-1 px-1 py-0.5 rounded text-[9px] font-mono font-bold bg-orange-100 text-orange-600 dark:bg-orange-500/20 dark:text-orange-400 align-text-top leading-none">
                            HTML
                        </span>
                    </h3>
                    <p class="text-[10px] text-slate-400 mt-1 font-mono truncate">{{ doc.path }}</p>
                </Link>
            </div>
        </div>
    </div>
</template>
