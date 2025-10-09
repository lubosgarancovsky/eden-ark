import { FC, ReactNode } from 'react';

type Props = {
    children: ReactNode
}

const Layout: FC<Props> = ({ children }) => {
 return (
      <div className='flex flex-col gap-8 items-center justify-center w-screen h-screen scroll-y-auto p-4 sm:p-8'>
          {children}
      </div>
 );
};

export default Layout;