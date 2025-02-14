'use client'

import { useLogin } from '@/hooks/public/clients/publicAPI'
import {
  Anchor,
  Avatar,
  Button,
  Center,
  Container,
  Group,
  Paper,
  PasswordInput,
  Space,
  Text,
  TextInput,
} from '@mantine/core'
import { useForm } from '@mantine/form'

function Login() {
  const form = useForm({
    initialValues: {
      email: '',
      password: '',
    },
    validate: {
      email: (value) =>
        /^\S+@\S+$/.test(value)
          ? null
          : 'メールアドレスの形式で入力してください',
      password: (value) =>
        value.length < 6 ? 'パスワードは6文字以上入力してください' : null,
    },
  })

  const { trigger, isMutating, error } = useLogin()

  const handleSubmit = async (values: { email: string; password: string }) => {
    await trigger(values)
  }

  return (
    <Center style={{ height: '100vh' }}>
      <Container style={{ width: 360, maxWidth: 360 }} my={40}>
        <Avatar
          size={100}
          src="/images/logo.png"
          style={{ margin: '0 auto' }}
        />
        <Text c="dimmed" size="sm" ta="center" mt={5}>
          アカウント未作成の方はこちら <br />
          <Anchor
            size="sm"
            component="button"
            onClick={() => (window.location.href = '/account/signup')}
          >
            新規アカウント作成
          </Anchor>
        </Text>

        <form onSubmit={form.onSubmit(handleSubmit)}>
          <Paper withBorder shadow="md" p={30} mt={30} radius="md">
            <Text ta="center" fw={700} size="xl">
              ログイン
            </Text>
            <Space h={5} />
            <TextInput
              label="メールアドレス"
              placeholder="example@gmail.com"
              required
              {...form.getInputProps('email')}
            />
            <PasswordInput
              label="パスワード"
              placeholder="password"
              required
              mt="md"
              {...form.getInputProps('password')}
            />

            <Group justify="space-between" mt="lg">
              <Anchor component="button" size="sm" style={{ margin: '0 auto' }}>
                パスワードをお忘れですか?
              </Anchor>
            </Group>

            {error && (
              <Text c="red" mt="md" ta="center">
                {error.message ?? 'ログインに失敗しました'}
              </Text>
            )}

            <Button fullWidth mt="xl" type="submit" disabled={isMutating}>
              {isMutating ? 'ログイン中...' : 'ログイン'}
            </Button>
          </Paper>
        </form>
      </Container>
    </Center>
  )
}

export default Login
