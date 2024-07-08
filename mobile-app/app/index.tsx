import { ThemedText } from "@/components/ThemedText";
import { useBabies } from "@/hooks/api/useBabies";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { Link } from "expo-router";
import { View } from "react-native";

const queryClient = new QueryClient();

export default function RootIndex() {
  return (
    <QueryClientProvider client={queryClient}>
      <View>
        <BabiesList />
      </View>
    </QueryClientProvider>
  );
}

function BabiesList() {
  const { isLoading, data: babies } = useBabies();
  if (isLoading || !babies) return <ThemedText>Loading</ThemedText>;

  return (
    <>
      {babies.map((baby) => (
        <Link href={`/${baby.id}`} key={baby.id}>
          <ThemedText type="link">Go to {baby.name}!</ThemedText>
        </Link>
      ))}
    </>
  );
}
