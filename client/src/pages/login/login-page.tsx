import {FC} from 'react';
import {Layout} from "../../components";
import {useLogin} from "../../hooks";
import {useSearchParams} from "react-router";

const LoginPage: FC = () => {
    const [urlSearchParams] = useSearchParams();
    const { onSubmit } = useLogin(urlSearchParams.get("returnTo") as string);


 return (
  <Layout>
      <div className="bg-neutral-800 p-8 rounded-xl">
          <form className='flex flex-col gap-4 min-w-72' onSubmit={onSubmit}>
              <div className='flex flex-col gap-1.5'>
                  <label htmlFor='email'>E-mail</label>
                  <input id='email' name="email" type="email" required/>
              </div>
              <div className='flex flex-col gap-1.5'>
                  <label htmlFor='name'>Password</label>
                  <input id='password' name="password" type="password" required/>
              </div>
              <button type='submit'>Login</button>
              <a className='text-sm text-center' href='#'>Forget password?</a>
          </form>
      </div>
  </Layout>
 );
};

export default LoginPage;