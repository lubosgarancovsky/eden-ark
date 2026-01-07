<script setup lang="ts">
type Props = {
   title: string;
 }

 const props = defineProps<Props>();
</script>

<template>
  <div class="flex items-center justify-center flex-col gap-10 w-96">
    <div class="flex items-center justify-center flex-col gap-4">
      <img src="/logo_48x48.svg" alt="eden-logo"/>
      <h3 class="font-normal">{{ props.title }}</h3>
    </div>
    <slot/>
  </div>
</template>
