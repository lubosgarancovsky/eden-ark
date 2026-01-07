<script setup lang="ts">
import {InputGroup, InputGroupAddon, InputGroupInput} from "@/components/ui/input-group";
import {Button} from "@/components/ui/button";
import { Eye, EyeOff } from 'lucide-vue-next'
import {type Component, ref} from "vue";

const isMasked = ref(true);

const togglePasswordMask = () => {
  isMasked.value = !isMasked.value;
}

interface Props {
  modelValue: string
  placeholder?: string
  icon?: Component
  type?: string
  name?: string
}

const props = defineProps<Props>()
const emit = defineEmits<{
  (e: 'update:modelValue', value: string): void
}>()

const updateValue = (e: Event) => {
  const target = e.target as HTMLInputElement
  emit('update:modelValue', target.value)
}

</script>

<template>
 <InputGroup>
  <InputGroupAddon v-if="props.icon">
    {{ props.icon }}
  </InputGroupAddon>
   <InputGroupInput :value="props.modelValue" @input="updateValue" :placeholder="props.placeholder" :type="isMasked ? 'password' : 'text'" :name="props.name"/>
  <InputGroupAddon align="inline-end">
    <Button @click="togglePasswordMask" size="icon-sm" variant="ghost" type="button" class="cursor-pointer">
      <Eye v-if="isMasked" />
      <EyeOff v-else/>
    </Button>
  </InputGroupAddon>
</InputGroup>
</template>
