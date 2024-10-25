interface ButtonProps {
  variant: "primary" | "secondary";
  children: React.ReactNode;
  type?: "button" | "submit" | "reset";
  onClick?: () => void;
  className?: string;
  disabled?: boolean;
}

export const Button = ({
  variant,
  type = "button",
  disabled,
  children,
  onClick,
  className,
}: ButtonProps) => {
  const baseStyles =
    "px-4 py-2 font-semibold rounded focus:outline-none focus:ring-2 focus:ring-offset-2";

  const variantStyles = {
    primary:
      "bg-indigo-600 text-white hover:bg-indigo-700 focus:ring-indigo-500",
    secondary:
      "bg-indigo-100 text-indigo-800 hover:bg-indigo-200 focus:ring-indigo-400",
  };

  return (
    <button
      type={type}
      className={`${baseStyles} ${variantStyles[variant]} ${className}`}
      disabled={disabled}
      onClick={onClick}
    >
      {children}
    </button>
  );
};
