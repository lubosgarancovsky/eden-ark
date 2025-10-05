import { FC, ReactNode } from 'react';

type Props = {
    children: ReactNode
}

const Layout: FC<Props> = ({ children }) => {
 return (
      <div className='flex items-center justify-center w-screen p-16'>
          {children}
      </div>
 );
};

export default Layout;