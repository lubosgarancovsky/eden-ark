<script setup lang="ts">

import {Form} from "vee-validate";
import {AdvancedInput, PasswordInput} from "@/components";
import {Input} from "@/components/ui/input";
import {useSignupForm} from "@/composables/useSignupForm.ts";
import {computed} from "vue";
import {Button} from "@/components/ui/button";
import {useRoute} from "vue-router";
import {ArrowLeft} from "lucide-vue-next";

const { isAvailable, firstName, lastName, username, email, password, passwordRepeat} = useSignupForm();

const usernameError = computed(() => {
  if (!isAvailable.value.isUsernameAvailable) {
    return "Username is already taken"
  }

  return username.error.value
})

const route = useRoute()
const errorMsg = route.query.error as string | undefined;
</script>

<template>
<Form class="flex flex-col gap-6 w-full max-w-sm" action="/oauth2/register" method="POST">
 <div class="flex flex-col gap-2">
   <AdvancedInput label="First name" :error="firstName.error.value">
     <Input v-model="firstName.value.value" name="firstName" />
   </AdvancedInput>

   <AdvancedInput label="Last name"  :error="lastName.error.value">
     <Input v-model="lastName.value.value" name="lastName"/>
   </AdvancedInput>

   <AdvancedInput label="Username" :error="usernameError">
     <Input v-model="username.value.value" name="username"/>
   </AdvancedInput>

   <AdvancedInput label="E-mail" :error="email.error.value || errorMsg">
     <Input v-model="email.value.value" name="email" />
   </AdvancedInput>

   <AdvancedInput label="Password" :error="password.error.value">
     <PasswordInput v-model="password.value.value" name="password" />
   </AdvancedInput>

   <AdvancedInput label="Repeat password" :error="passwordRepeat.error.value">
     <PasswordInput v-model="passwordRepeat.value.value" name="passwordRepeat"/>
   </AdvancedInput>
 </div>

  <div class="flex flex-col gap-4 w-full flex-1">
    <Button type="submit" class="w-full cursor-pointer">Sign up</Button>
    <Button variant="secondary" class="w-full" type="button" asChild>
      <a href="/login">
        <ArrowLeft/>
        Back to login
      </a>
    </Button>
  </div>
</Form>
</template>
