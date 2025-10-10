import {FC, DetailedHTMLProps, ButtonHTMLAttributes} from 'react';
import {cn} from "../../../lib";

type Props = DetailedHTMLProps<ButtonHTMLAttributes<HTMLButtonElement>, HTMLButtonElement> & {
    variant?: 'primary' | 'secondary'
}

const Button: FC<Props> = ({children, variant = 'primary', className, ...props}) => {
 return (
  <button className={cn("text-sm px-4 py-2 rounded-full cursor-pointer font-medium", {
      "bg-blue-700 font-medium hover:bg-blue-600 disabled:bg-black/20 disabled:text-neutral-400": variant === 'primary',
      "bg-transparent hover:bg-black/20 text-neutral-100": variant === 'secondary'
  }, className)} {...props}>
      {children}
  </button>
 );
};

export default Button;