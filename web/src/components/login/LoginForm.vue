<script setup lang="ts">
import { Form } from 'vee-validate'
import {Input} from '@/components/ui/input'
import {Button} from "@/components/ui/button";
import {AdvancedInput, PasswordInput} from "@/components";
import {useLoginForm} from "@/composables";
import { useRoute } from 'vue-router'

const route = useRoute()
const errorMsg = route.query.error as string | undefined;
const returnTo = route.query.returnTo as string | undefined;

const { username, password, emailError, passwordError} = useLoginForm();
</script>

<template>
  <Form id="registerForm" action="/oauth2/login" method="POST" class="flex flex-col gap-2 w-full max-w-sm">
    <div class="flex flex-col gap-4">
      <input type="hidden" name="returnTo" :value="returnTo" />
      <AdvancedInput label="Username or primary e-mail" :error="errorMsg || emailError">
        <Input v-model="username" name="username" type="text" />
      </AdvancedInput>
      <AdvancedInput label="Password" :error="passwordError">
        <PasswordInput v-model="password" name="password" type="password"/>
      </AdvancedInput>
    </div>
    <Button variant="link" class='p-0 ml-auto' type="button" asChild>
      <a href="/forgot-password">Forgot your password?</a>
    </Button>

    <div class="flex flex-col items-center justify-center gap-2 w-full flex-1">
      <Button type="submit" class="w-full cursor-pointer">Sign in</Button>
      <div class="relative flex flex-col items-center justify-center text-sm text-muted-foreground w-full">
        <div class="w-full h-px border-b absolute left-0 top-1/2 -z-1"/>
        <div class='px-4 py-1 bg-background'>Don't have an account yet ?</div>
      </div>
      <Button variant="secondary" class="w-full" type="button" asChild>
        <a href="/signup">Create an account</a>
      </Button>
    </div>
  </Form>
</template>