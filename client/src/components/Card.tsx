import React from "react";


export const Card = ({
    children,
  }: {
    children: React.ReactNode
  }) => {
    return (
        <div className="w-11/12 mx-1 mb-4 p-8 bg-pf-gray-900 rounded-xl">
            {children}
        </div>
    );
}