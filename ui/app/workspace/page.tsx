"use client";

import { useRouter } from "next/navigation";
import { useEffect } from "react";

export default function WorkspacePage() {
	const router = useRouter();
	
	useEffect(() => {
		router.replace("/workspace/dashboard");
	}, [router]);
	
	return null;
}
