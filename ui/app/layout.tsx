import "./globals.css"

export default function RootLayout({ children }: { children: React.ReactNode }) {
	return (
		<html lang="en" suppressHydrationWarning>
			<head>
				<link rel="dns-prefetch" href="https://getbifrost.ai" />
				<link rel="preconnect" href="https://getbifrost.ai" />
			</head>
			<body className="font-sans antialiased">
				{children}
			</body>
		</html>
	)
}
