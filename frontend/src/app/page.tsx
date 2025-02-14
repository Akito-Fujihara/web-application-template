'use client'

import { Avatar, Button, Center, Container, Stack, Title } from '@mantine/core'

export default function Home() {
  return (
    <Center style={{ height: '100vh' }}>
      <Container
        style={{
          width: '400px', // 幅を固定
          border: '2px solid #696969', // 外枠を追加
          borderRadius: '8px', // 角を丸くする（オプション）
          padding: '20px', // 内側の余白
          backgroundColor: 'white', // 背景色を追加（オプション）
        }}
      >
        <Stack align="center">
          <Avatar size={200} src="images/logo.png" />
          <Title order={2}>Web Application Template</Title>
          <Button
            className="mt-4"
            size="lg"
            variant="filled"
            color="blue"
            fullWidth
            onClick={() => (window.location.href = '/account/login')}
          >
            ログイン
          </Button>
          <Button
            size="lg"
            variant="outline"
            color="blue"
            fullWidth
            onClick={() => (window.location.href = '/account/signup')}
          >
            新規アカウント作成
          </Button>
        </Stack>
      </Container>
    </Center>
  )
}
