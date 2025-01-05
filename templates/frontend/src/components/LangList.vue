<script setup>
import { getLanguages, addLanguage, deleteLanguage } from '../fetch/lang';
import { onMounted, ref } from 'vue';

const langs = ref([])
onMounted(async () => {
    const result = await getLanguages()
    if (result.error) {
        console.error('Error:', result.message)
    } else {
        console.log('Languages:', result.data)
        langs.value = result.data
    }
})

const langName = ref('')
const releaseYear = ref(0)

const newLanguage = async () => {
    const result = await addLanguage(langName.value, releaseYear.value)
    console.log(result)
    if (result.error) {
        console.error('Error:', result.message)
    } else {
        console.log('New language id:', result.languageId)
        let newLang = {
            id: result.languageId,
            name: langName,
            release_year: releaseYear
        }
        langs.value = [...langs.value, newLang]
    }
}

const removeLanguage = async (langId) => {
    const result = await deleteLanguage(langId)
    if (result.error) {
        console.error('Error:', result.message)
    } else {
        console.log('Language deleted')
        const updateLangsArr = langs.value.filter(lang => lang.id !== langId)
        langs.value = updateLangsArr
    }
}
</script>

<template>
    <form @submit.prevent="newLanguage">
        <label>Add language</label> <input type="text" placeholder="Golang" v-model="langName">
        <label>Release year</label><input type="number" placeholder="2009" v-model="releaseYear">
        <button type="submit" id="add-btn">Add</button>
    </form>
    <div id="list-container">
        <div v-for="lang in langs" id="lang-info">
            <p>{{ lang.name }}</p>
            <p>Release: {{ lang.release_year }}</p>
            <button @click="removeLanguage(lang.id)">Delete</button>
        </div>
    </div>
</template>

<style scoped>
#lang-info {
    background-color: rgb(32, 32, 100);
    border-radius: 0.5em;
    padding: 0.5em;
    margin-top: 8px;

    p {
        margin: 0;
        padding: 0;
    }

    &:hover {
        background-color: rgb(45, 45, 131);
    }
}

button {
    cursor: pointer;
}

#add-btn {
    padding: 0.3em;
}
</style>