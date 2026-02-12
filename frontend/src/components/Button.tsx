import React from 'react';

interface ButtonProps extends React.ButtonHTMLAttributes<HTMLButtonElement> {
  variant?: 'primary' | 'secondary' | 'outline';
  isLoading?: boolean;
  loadingText?: string;
  children: React.ReactNode;
}

export default function Button({
  variant = 'primary',
  isLoading = false,
  loadingText,
  children,
  className = '',
  disabled,
  ...props
}: ButtonProps) {
  const baseStyles = 'w-full py-3 rounded-lg font-medium transition-colors cursor-pointer disabled:opacity-50 disabled:cursor-not-allowed';
  
  const variantStyles = {
    primary: 'bg-teal-600 text-white hover:bg-teal-700',
    secondary: 'bg-gray-600 text-white hover:bg-gray-700',
    outline: 'border-2 border-teal-600 text-teal-600 hover:bg-teal-50',
  };

  return (
    <button
      className={`${baseStyles} ${variantStyles[variant]} ${className}`}
      disabled={disabled || isLoading}
      {...props}
    >
      {isLoading ? (loadingText || 'Carregando...') : children}
    </button>
  );
}

