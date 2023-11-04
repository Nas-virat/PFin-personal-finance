import React from "react";


export const Card = ({
    children,
  }: {
    children: React.ReactNode
  }) => {
    return (
        <div className="w-full mx-1 mb-4 p-9 bg-pf-gray-900 rounded-xl">
            {children}
        </div>
    );
}