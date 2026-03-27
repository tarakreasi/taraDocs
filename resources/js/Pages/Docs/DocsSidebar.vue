<script setup>
import { computed, ref } from 'vue';
import axios from 'axios';
import { useTheme } from '@/composables/useTheme';

const props = defineProps({
    allDocs: { type: Array, default: () => [] },
    selectedCategory: { type: String, default: null },
});

const emit = defineEmits(['select-category']);

const { isDark, toggleTheme } = useTheme();

const categories = computed(() => {
    return [...new Set(props.allDocs.map(doc => doc.category))];
});

const isCreatingFolder = ref(false);

const createNewFolder = async () => {
    const folderName = prompt("Enter new folder name:");
    if (!folderName) return;

    // Use "New Categories" as default parent if nothing is selected, else try to use selected as parent
    let parentPath = "";
    if (props.selectedCategory && props.selectedCategory !== "Uncategorized") {
        parentPath = props.selectedCategory.replace(/ \/ /g, "/"); // Convert "A / B" to "A/B"
    }

    try {
        isCreatingFolder.value = true;
        const response = await axios.post('/docs/api/create-folder', {
            parentPath: parentPath,
            folderName: folderName,
        });
        
        if (response.data.success) {
            emit('category-created');
        }
    } catch (error) {
        alert(error.response?.data?.error || "Failed to create folder");
    } finally {
        isCreatingFolder.value = false;
    }
};
</script>

<template>
    <aside class="w-[240px] flex flex-col bg-slate-50/50 dark:bg-[#0b1121] border-r border-slate-200 dark:border-white/5 shrink-0 transition-all duration-300">

        <!-- Brand -->
        <div class="h-16 flex items-center px-4 gap-3 shrink-0 border-b border-slate-200 dark:border-white/5">
            <div class="size-8 rounded-lg bg-indigo-500 flex items-center justify-center text-white shadow-sm">
                <span class="material-symbols-outlined text-[20px]">description</span>
            </div>
            <span class="font-sans font-bold text-base tracking-tight text-slate-800 dark:text-white">Docs</span>
            
            <!-- Theme toggle -->
            <button 
                @click="toggleTheme" 
                class="ml-auto size-8 rounded-lg flex items-center justify-center text-slate-400 hover:text-indigo-500 hover:bg-slate-100 dark:hover:bg-white/5 transition-colors"
                :title="isDark ? 'Switch to Light' : 'Switch to Dark'"
            >
                <span class="material-symbols-outlined text-[20px]">{{ isDark ? 'light_mode' : 'dark_mode' }}</span>
            </button>
        </div>

        <!-- Navigation -->
        <nav class="flex-1 px-3 py-3 space-y-0.5 overflow-y-auto">
            <!-- All Docs -->
            <button
                @click="emit('select-category', null)"
                class="w-full flex items-center gap-3 px-3 py-2 rounded-lg transition-all text-sm text-left"
                :class="selectedCategory === null
                    ? 'bg-white dark:bg-white/5 text-slate-900 dark:text-white font-medium shadow-sm ring-1 ring-black/5 dark:ring-white/5'
                    : 'text-slate-500 dark:text-slate-400 hover:bg-slate-100 dark:hover:bg-white/5 hover:text-slate-900 dark:hover:text-slate-200'"
            >
                <span class="material-symbols-outlined text-[20px]">library_books</span>
                <span class="flex-1">All Docs</span>
                <span class="font-mono text-[10px] text-slate-400">{{ allDocs.length }}</span>
            </button>

            <!-- Categories Header -->
            <div class="px-3 mt-6 mb-2 flex items-center justify-between">
                <span class="text-[10px] font-bold text-slate-400 dark:text-slate-500 uppercase tracking-widest">Categories</span>
                <button 
                    @click="createNewFolder"
                    :disabled="isCreatingFolder"
                    class="size-5 rounded flex items-center justify-center text-slate-400 hover:text-indigo-500 hover:bg-slate-200 dark:hover:bg-white/10 transition-colors disabled:opacity-50"
                    title="New Folder"
                >
                    <span class="material-symbols-outlined text-[14px]">create_new_folder</span>
                </button>
            </div>

            <!-- Empty State -->
            <div v-if="categories.length === 0" class="px-3 py-2 text-xs text-slate-400 italic">
                No categories found
            </div>

            <!-- Category Buttons -->
            <button
                v-for="category in categories"
                :key="category"
                @click="emit('select-category', category)"
                class="w-full flex items-center gap-3 px-3 py-2 rounded-lg transition-all text-sm text-left"
                :class="selectedCategory === category
                    ? 'bg-white dark:bg-white/10 ring-1 ring-indigo-500 text-slate-900 dark:text-white font-medium shadow-sm'
                    : 'text-slate-500 dark:text-slate-400 hover:bg-slate-100 dark:hover:bg-white/5 hover:text-slate-700 dark:hover:text-slate-300'"
            >
                <span class="material-symbols-outlined text-[18px] shrink-0" :class="selectedCategory === category ? 'text-indigo-500' : 'text-slate-400 opacity-80'">folder</span>
                <span class="flex-1 truncate">{{ category }}</span>
            </button>
        </nav>
    </aside>
</template>
