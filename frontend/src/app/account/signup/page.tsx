'use client'

import { useSignUp } from '@/hooks/public/clients/publicAPI'
import {
  Anchor,
  Avatar,
  Button,
  Center,
  Container,
  Paper,
  PasswordInput,
  Space,
  Text,
  TextInput,
} from '@mantine/core'
import { useForm } from '@mantine/form'

function SignUp() {
  const form = useForm({
    initialValues: {
      username: '',
      email: '',
      password: '',
      passwordConfirmation: '',
    },
    validate: {
      username: (value) =>
        value.length < 2 ? 'お名前は2文字以上入力してください' : null,
      email: (value) =>
        /^\S+@\S+$/.test(value)
          ? null
          : 'メールアドレスの形式で入力してください',
      password: (value) =>
        value.length < 6 ? 'パスワードは6文字以上入力してください' : null,
      passwordConfirmation: (value, values) =>
        value === values.password ? null : 'パスワードが一致しません',
    },
  })

  const { trigger, isMutating, error } = useSignUp()

  const handleSubmit = async (values: {
    username: string
    email: string
    password: string
  }) => {
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
          アカウント作成済みの方はこちら <br />
          <Anchor
            size="sm"
            component="button"
            onClick={() => (window.location.href = '/account/login')}
          >
            ログイン
          </Anchor>
        </Text>

        <form onSubmit={form.onSubmit(handleSubmit)}>
          <Paper withBorder shadow="md" p={30} mt={30} radius="md">
            <Text ta="center" fw={700} size="xl">
              新規アカウント作成
            </Text>
            <Space h={5} />
            <TextInput
              label="お名前"
              placeholder="山田 太郎"
              required
              mt="md"
              {...form.getInputProps('username')}
            />
            <TextInput
              label="メールアドレス"
              placeholder="example@gmail.com"
              required
              mt="md"
              {...form.getInputProps('email')}
            />
            <PasswordInput
              label="パスワード"
              placeholder="password"
              required
              mt="md"
              {...form.getInputProps('password')}
            />
            <PasswordInput
              label="パスワード（確認）"
              placeholder="password"
              required
              mt="md"
              {...form.getInputProps('passwordConfirmation')}
            />

            {error && (
              <Text c="red" mt="md" ta="center">
                {error.message ?? 'アカウント作成に失敗しました'}
              </Text>
            )}

            <Button fullWidth mt="xl" type="submit" disabled={isMutating}>
              {isMutating ? 'アカウント作成中...' : 'アカウント作成'}
            </Button>
          </Paper>
        </form>
      </Container>
    </Center>
  )
}

export default SignUp
