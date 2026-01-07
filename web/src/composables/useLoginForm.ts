import {useForm, useField} from 'vee-validate'
import {toTypedSchema} from "@vee-validate/zod";
import { z } from 'zod'

const loginSchema = toTypedSchema(
    z.object({
        username: z.string().min(1, "Username or e-mail is required"),
        password: z.string().min(1, 'Password is required')
    })
)


export function useLoginForm() {
    const { resetForm } = useForm({
        validationSchema: loginSchema,
        validateOnMount: false,
        initialValues: {
            username: '',
            password: '',
        },
    })

    const { value: username, errorMessage: emailError } = useField<string>('username')
    const { value: password, errorMessage: passwordError } = useField<string>('password')

    return { username, emailError, password, passwordError, resetForm }
}