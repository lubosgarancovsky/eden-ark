import {FC, ReactNode} from 'react';
import {cn} from "../../../lib";

type Props = {
    title: string;
    children: ReactNode
    variant?: 'error' | 'success'
}

const ErrorBanner: FC<Props> = ({ title, children, variant = 'error' }) => {
 return (
  <div className={cn("border-l-2 flex flex-col gap-1 p-4", {"border-red-500 bg-red-500/10 text-red-400": variant === 'error', "border-green-500 bg-green-500/10 text-green-400" : variant === 'success'})}>
      <div className='font-medium '>{title}</div>
      <div>{children}</div>
  </div>
 );
};

export default ErrorBanner;