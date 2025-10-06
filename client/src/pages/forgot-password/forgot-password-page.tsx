import  { FC } from 'react';
import {Button, ErrorBanner, Layout} from "../../components";
import Card from "../../components/core/card/card.tsx";
import {useResetPassword} from "../../hooks";

const ForgotPasswordPage: FC = () => {
   const { error, email, state, onSubmit, resend } = useResetPassword();

 return (
  <Layout>
      <Card title="Reset password" description="Enter your e-mail to reset password">
          {
              error && <ErrorBanner title="An error has occured">{error.message}</ErrorBanner>
          }
          {state === 'success' && <div className='flex flex-col gap-4'>
              <p>An e-mail with a verification link was sent to your e-mail address:</p>
              <div className='p-4 bg-neutral-800 rounded-lg text-neutral-300 text-center'>{email}</div>

              <div className="flex items-center justify-between gap-8">
                  <a href="/login">
                      <Button type='button' variant="secondary">Back to login</Button>
                  </a>
                  <Button variant="primary" onClick={resend}>Resend e-mail</Button>
              </div>
          </div>}
          {state !== 'success' && <form className='flex flex-col gap-8 min-w-72' onSubmit={onSubmit}>
              <div className="flex flex-col gap-4">
                  <div className='flex flex-col gap-1.5'>
                      <label htmlFor='email'>E-mail</label>
                      <input id='email' name="email" type="email" required/>
                  </div>
              </div>

              <p className="text-neutral-300">
                  A verification link will be sent to your e-mail
              </p>

              <div className="flex items-center justify-between gap-8">
                  <a href="/login">
                      <Button type='button' variant="secondary">Back to login</Button>
                  </a>
                  <Button variant="primary" disabled={state === 'pending'}>Send verification e-mail</Button>
              </div>
          </form>}
      </Card>
  </Layout>
 );
};

export default ForgotPasswordPage;