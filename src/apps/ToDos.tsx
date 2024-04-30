"use client";

import React from "react";
import { CheckboxWithLabel } from "~/components/ui/checkbox";
import { api } from "~/trpc/react";

const ToDos: React.FC = () => {
  const {
    data: dueToDos,
    isPending,
    isError,
  } = api.todo.due.useQuery(undefined, {
    refetchInterval: 1000 * 60 * 60 * 15, // 15 minutes
  });

  if (isPending || isError) {
    return null;
  }

  return (
    <div className="flex flex-col gap-2">
      {dueToDos.map((todo) => (
        <CheckboxWithLabel
          variant={todo.isOverdue ? "danger" : undefined}
          key={todo.id}
        >
          {todo.content}
        </CheckboxWithLabel>
      ))}

      {dueToDos.length === 0 ? <div>No tasks for today :)</div> : null}
    </div>
  );
};

export default ToDos;
