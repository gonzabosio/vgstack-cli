<script setup></script>

<template>
    <div>
        <div id="leading-container">
            <p>Go+Vue App</p>
            <p>Based on <span><a href="https://vite.dev" target="_blank" id="vite-link">Vite</a></span>⚡ scaffolding</p>
            <RouterLink to="/languages" id="nav-lang">Languages List</RouterLink>
        </div>
    </div>
</template>
<style scoped>
#leading-container {
    border: 16px dashed;
    border-radius: 1em;
    padding: 2em;
    transition: 0.5s;
    border-left-color: rgb(98, 159, 163);
    border-top-color: rgb(98, 159, 163);
    border-right-color: rgb(34, 114, 45);
    border-bottom-color: rgb(34, 114, 45);

    &:hover {
        transform: scale(1.05);
        border-right-color: rgb(98, 159, 163);
        border-bottom-color: rgb(98, 159, 163);
        border-left-color: rgb(34, 114, 45);
        border-top-color: rgb(34, 114, 45);
    }
}

#vite-link {
    color: rgb(168, 111, 221);
    text-decoration: none;
    transition: 0.3s;

    &:hover {
        text-shadow: rgb(198, 166, 228) 1px 0 15px;
        ;
    }
}

#nav-lang {
    display: block;
    transition: 0.3s;
    text-decoration: none;
    color: rgba(255, 255, 255, 0.85);
    padding: 0.8em;
    border-radius: 1em;
    background-color: rgb(39, 74, 139);

    &:hover {
        box-shadow: 0px 8px 8px 8px rgba(39, 74, 139, 0.226);
    }

}
</style>