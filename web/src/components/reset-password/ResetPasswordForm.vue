<script setup lang="ts">
import { useRoute, useRouter } from "vue-router";
import { Form, type SubmissionHandler, useField, useForm } from "vee-validate";
import { AdvancedInput, ErrorBanner, PasswordInput } from "@/components";
import { toTypedSchema } from "@vee-validate/zod";
import { string, object } from 'zod';
import { useCountdown } from "@/composables";
import { Button } from "@/components/ui/button";
import { useMutation } from "@tanstack/vue-query";
import  { type ApiError, changePassword } from "@/lib/api";

const schema = toTypedSchema(
    object({
        password: string().min(8, "Password must be at least 8 characters long").max(100),
        passwordRepeat: string(),
    }).refine((data) => data.password === data.passwordRepeat, {
        message: "Passwords must match",
        path: ["passwordRepeat"],
    })
)

type ResetPasswordFormType = {  password: string, passwordRepeat: string, token: string}

const route = useRoute();
const router = useRouter();
const token = route.query.token as string;
const expiresAt = route.query.expiresAt as string ?? new Date().toISOString();

const { percent, formattedTime } = useCountdown(expiresAt);
const mutation = useMutation<void, ApiError, { password: string, token: string}>({
    mutationFn: ({ password, token }) => changePassword(password, token)
})

const { handleSubmit } = useForm<ResetPasswordFormType>({
    validationSchema: schema,
    initialValues: {
        password: '',
        passwordRepeat: ''
    }
})

const { value: password, errorMessage: passwordError } = useField<string>('password')
const { value: passwordRepeat, errorMessage: passwordRepeatError } = useField<string>('passwordRepeat')

const onSubmit: SubmissionHandler<any> = handleSubmit((values) => {
    mutation.mutate({ token, password: values.password}, {
        onSuccess: () => {
            router.push('/login')
        }
    });
})

</script>

<template>
  <Form class="flex flex-col gap-4 w-full max-w-sm" @submit="onSubmit">
      <ErrorBanner v-if="mutation.isError.value">
          {{ mutation.error.value?.message ?? "An unexpected error occurred. Please try again later." }}
      </ErrorBanner>
      <div class="flex flex-col gap-4 w-full">
          <AdvancedInput label="New password" :error="passwordError">
              <PasswordInput v-model="password" name="password" />
          </AdvancedInput>
          <AdvancedInput label="Repeat password" :error="passwordRepeatError">
              <PasswordInput v-model="passwordRepeat" name="passwordRepeat" />
          </AdvancedInput>
      </div>

      <div class="w-full flex items-center gap-2 text-sm text-muted-foreground">
          <div class="text-nowrap">Expires in</div>
          <div class='bg-secondary w-full rounded-full h-1 overflow-hidden'>
              <div class="transition-all duration-200 bg-primary h-full" :style="{ width: percent + '%'}"/>
          </div>
          <div>{{ formattedTime }}</div>
      </div>

      <Button type="submit" class="w-full cursor-pointer" :disabled="percent <= 0">Change password</Button>
  </Form>
</template>