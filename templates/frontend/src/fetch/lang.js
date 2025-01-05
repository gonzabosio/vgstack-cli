const backurl = import.meta.env.VITE_BACKEND_URL

export const getLanguages = async () => {
    try {
        const resp = await fetch(backurl + '/language')
        const payload = await resp.json()
        // console.log(payload)
        if (!resp.ok) {
            console.error(payload.error_dets)
            return { error: true, message: payload.message }
        }
        return { error: false, data: payload.languages }
    } catch (error) {
        console.error(error.message)
        return { error: true, message: error.message }
    }
}

export const addLanguage = async (langName, releaseYear) => {
    try {
        const resp = await fetch(backurl + '/language', {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ name: langName, release_year: releaseYear })
        })
        const payload = await resp.json()
        // console.log(payload)
        if (!resp.ok) {
            console.error(payload.error_dets)
            return { error: true, message: payload.message }
        }
        return { error: false, languageId: payload.lang_id }
    } catch (error) {
        console.error(error.message)
        return { error: true, message: error.message }
    }
}

export const deleteLanguage = async (langId) => {
    try {
        const resp = await fetch(backurl + `/language/${langId}`, {
            method: "DELETE",
        })
        const payload = await resp.json()
        // console.log(payload)
        if (!resp.ok) {
            console.error(payload.error_dets)
            return { error: true, message: payload.message }
        }
        return { error: false, message: payload.message }
    } catch (error) {
        console.error(error.message)
        return { error: true, message: error.message }
    }
}