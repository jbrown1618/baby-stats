import { ThemedText } from "@/components/ThemedText";
import { useBabies } from "@/hooks/api/useBabies";
import { Link } from "expo-router";
import { View } from "react-native";

export default function RootIndex() {
  const { isLoading, data: babies } = useBabies();
  if (isLoading || !babies) return <ThemedText>Loading</ThemedText>;

  return (
    <View>
      {babies.map((baby) => (
        <Link href={`/${baby.id}`} key={baby.id}>
          <ThemedText type="link">Go to {baby.name}!</ThemedText>
        </Link>
      ))}
    </View>
  );
}
