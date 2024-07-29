import React from "react";

export const PageContainer = ({ children }: { children: React.ReactNode }) => {
  return (
    <div className="w-full mt-5 px-5 flex flex-col justify-center items-center">
      {children}
    </div>
  );
};

