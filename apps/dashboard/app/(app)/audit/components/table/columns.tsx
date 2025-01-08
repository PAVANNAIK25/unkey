import { TimestampInfo } from "@/components/timestamp-info";
import { Badge } from "@/components/ui/badge";
import type { Column } from "@/components/virtual-table/types";
import { cn } from "@unkey/ui/src/lib/utils";
import { FunctionSquare, KeySquare } from "lucide-react";
import type { Data } from "./types";
import { getEventType } from "./utils";

export const columns: Column<Data>[] = [
  {
    key: "time",
    header: "Time",
    headerClassName: "pl-3",
    width: "150px",
    render: (log) => (
      <div className="flex items-center gap-3 px-2">
        <TimestampInfo
          value={log.auditLog.time}
          className="font-mono group-hover:underline decoration-dotted"
        />
      </div>
    ),
  },
  {
    key: "actor",
    header: "Actor",
    headerClassName: "pl-3",
    width: "7%",
    render: (log) => (
      <div className="flex items-center gap-3 px-2">
        {log.auditLog.actor.type === "user" && log.user ? (
          <div className="flex items-center w-full gap-2 max-sm:m-0 max-sm:gap-1 max-sm:text-xs">
            <span className="text-xs whitespace-nowrap">{`${log.user.firstName ?? ""} ${
              log.user.lastName ?? ""
            }`}</span>
          </div>
        ) : log.auditLog.actor.type === "key" ? (
          <div className="flex items-center w-full gap-2 max-sm:m-0 max-sm:gap-1 max-sm:text-xs">
            <KeySquare className="w-4 h-4" />
            <span className="font-mono text-xs">{log.auditLog.actor.id}</span>
          </div>
        ) : (
          <div className="flex items-center w-full gap-2 max-sm:m-0 max-sm:gap-1 max-sm:text-xs">
            <FunctionSquare className="w-4 h-4" />
            <span className="font-mono text-xs">{log.auditLog.actor.id}</span>
          </div>
        )}
      </div>
    ),
  },
  {
    key: "action",
    header: "Action",
    headerClassName: "pl-3",
    width: "7%",
    render: (log) => {
      const eventType = getEventType(log.auditLog.event);
      const badgeClassName = cn("font-mono capitalize", {
        "bg-error-3 text-error-11 hover:bg-error-4": eventType === "delete",
        "bg-warning-3 text-warning-11 hover:bg-warning-4": eventType === "update",
        "bg-success-3 text-success-11 hover:bg-success-4": eventType === "create",
        "bg-accent-3 text-accent-11 hover:bg-accent-4": eventType === "other",
      });
      return (
        <div className="flex items-center gap-3 px-2">
          <Badge className={badgeClassName}>{eventType}</Badge>
        </div>
      );
    },
  },
  {
    key: "event",
    header: "Event",
    headerClassName: "pl-2",
    width: "20%",
    render: (log) => (
      <div className="flex items-center gap-2 text-current font-mono text-xs px-2">
        <span>{log.auditLog.event}</span>
      </div>
    ),
  },
  {
    key: "event-description",
    header: "Description",
    headerClassName: "pl-1",
    width: "auto",
    render: (log) => (
      <div className="text-current font-mono px-2 text-xs">{log.auditLog.description}</div>
    ),
  },
];
