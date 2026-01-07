<script setup lang="ts">
import {Label} from "@/components/ui/label";


interface Props {
  label?: string;
  error?: string;
  optional?: boolean;
}

const props = defineProps<Props>()
</script>

<template>
<div class='flex flex-col gap-1' :class="{'advanced-input--error': props.error}">
  <Label v-if="props.label" class='mb-1'>
    <span>{{ props.label }}</span>
    <span v-if="props.optional" class='text-muted-foreground text-sm'>(Optional)</span>
  </Label>
  <slot />
  <div v-if="props.error" class="text-sm text-red-500">{{ props.error }}</div>
</div>
</template>
