import {FC, DetailedHTMLProps, ButtonHTMLAttributes} from 'react';
import {cn} from "../../../lib";

type Props = DetailedHTMLProps<ButtonHTMLAttributes<HTMLButtonElement>, HTMLButtonElement> & {
    variant?: 'primary' | 'secondary'
}

const Button: FC<Props> = ({children, variant = 'primary', className, ...props}) => {
 return (
  <button className={cn("px-5 py-3 rounded-full cursor-pointer font-medium", {
      "bg-blue-700 font-medium hover:bg-blue-600 disabled:bg-neutral-700/40 disabled:text-neutral-400": variant === 'primary',
      "bg-neutral-900 hover:bg-neutral-400/15 text-neutral-100": variant === 'secondary'
  }, className)} {...props}>
      {children}
  </button>
 );
};

export default Button;